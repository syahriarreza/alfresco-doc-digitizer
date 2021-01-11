// SETTING VARIABLES
// Note: FILE_DATE -> text, format: yyyy/mm/dd, ex: 2011/04/20
var prefix = "efilling";
var typeQname = "efilling:efillingType";
var splitter1 = ","; var splitter2 = ":-"; // ex: MetadataName1:-MetaValue1,MetadataName2:-MetaValue2,... >> "," -> 1   ":-" -> 2
var mandatoryMetadata = ["CLIENT", "EXT_REF_ID", "FILE_DATE"]; // Mandatory Metadata (metadata that used for creating folders)
var metadata2;

// Check Missing Metadata, Return True: metadata is missing, False: metadata is available
function checkMissingMetadata(metadataNameArray, metadataName){
	if(metadataNameArray.indexOf(metadataName) < 0)
	{
		result = "fail";
		info = "Uploading fail due to undefined "+metadataName;
		return true;
	}
	return false;
}

// Ensure Folders in path exist, if not, it will be created
function ensureFolderExistByPath(path){
	var folderNameArray = path.split("/"); // Get Folder names from path

	var destFoldNode;
	for(var i = 0; i<folderNameArray.length; i++){
		if(i == 0){
			// Location is in Root of Repository
			var checkFoldNode = companyhome.childByNamePath(folderNameArray[i]);
			if(checkFoldNode != null){destFoldNode = checkFoldNode;}
			else{destFoldNode = companyhome.createFolder(folderNameArray[i]);}
		}else{
			// Location is inside of a folder
			var checkFoldNode = destFoldNode.childByNamePath(folderNameArray[i]);
			if(checkFoldNode != null){destFoldNode = checkFoldNode;}
			else{destFoldNode = destFoldNode.createFolder(folderNameArray[i]);}
		}
	}
	return destFoldNode;
}

// Splitting Metadata into Array; it will return an Array containing 2 arrays
function splitMetadata(metadata){
	var metadataArrays = new Array();
	var metadata2DArray = new Array(); // #1 Col: Metadata Name ; #2 Col: Metadata Value
	var metadataCol1DArray = new Array(); // Array of metadata name for searching
	var rows = metadata.split(splitter1);
	for(var i=0; i<rows.length; i++){
		logger.log("LOG INFO ===== rows["+i+"]: "+rows[i]);
		var row = rows[i].split(splitter2);
		if(row[1] == null) row[1] = "";
		metadata2DArray.push(row);
		metadataCol1DArray.push(row[0]);
	}
	metadataArrays.push(metadata2DArray); metadataArrays.push(metadataCol1DArray);
	return metadataArrays;
}


// Main Function
function main() {
	var filename = null;
	var content = null;
	var metadata = String(args.metadata);
	var metadata2 = String(args.metadata2);

	// locate file attributes
	// formdata is data which got from the uploadSPT.get.html.ftl and has many fields
	// GET file attributes and other data
	var formFields = formdata.fields;
	for(var i=0; i<formFields.length; i++){
		logger.log("LOG INFO FORMFIELD = "+formFields[i].name+": "+formFields[i].value);
	}
	for each (field in formdata.fields)
	{
		switch (String(field.name)){
			case "file":
				if (field.isFile)
				{
					filename = field.filename;
					content = field.content;
				}
				break;
			case "metadata":
				metadata = field.value; logger.log("LOG INFO ===== metadata: "+metadata);
				break;
			case "metadata2":
				logger.log("LOG INFO ===== field: "+field);
				metadata2 = field.value; logger.log("LOG INFO ===== metadata2: "+metadata2);
				break;
			default:
				logger.log("LOG WARN ===== Variable \""+field.name+"\": \""+field.value+"\" is not received by the webscript");
				break;
		}
	}

	// Logger Data
	logger.log("LOG INFO ===== filename: " + filename);
	logger.log("LOG INFO ===== content: " + content);
	metadata = String(metadata); // Make sure metadata is String, so later can be splitted

	// CHECKING file name and content
	if (filename == undefined || content == undefined)
	{
		// UNDEFINED FILE
		result = "fail";
		info = "Uploading fail due to undefined file";
		return;
	}

	// Splitting Metadata into Array
	var metadataArrays = splitMetadata(metadata);
	var metadata2DArray = metadataArrays[0]; // #1 Col: Metadata Name ; #2 Col: Metadata Value
	var metadataCol1DArray = metadataArrays[1]; // Array of metadata name for searching

	// CHECKING Mandatory Metadata (metadata that used for creating folders)
	for(var i=0; i<mandatoryMetadata.length; i++){
		if(checkMissingMetadata(metadataCol1DArray, mandatoryMetadata[i])) return;
	}
	
	// Generating Path from Metadata; pathPattern = "Kontrak/CLIENT/Year/Month/EXT_REF_ID/DOC_TYPE";
	var CLIENT = metadata2DArray[metadataCol1DArray.indexOf("CLIENT")][1]; CLIENT = String(CLIENT).replace("/", "-");
	var EXT_REF_ID = metadata2DArray[metadataCol1DArray.indexOf("EXT_REF_ID")][1]; EXT_REF_ID = String(EXT_REF_ID).replace("/", "-");
	var dateString = metadata2DArray[metadataCol1DArray.indexOf("FILE_DATE")][1];
	// check FILE_DATE
	var currDate = new Date(dateString);
	if(currDate == null || currDate == undefined)
	{
		result = "fail";
		info = "Uploading fail due to incorrect FILE_DATE which is \""+dateString+"\"";
		return;
	}
	// Generate Year and Month
	var yearStr = currDate.getFullYear();
	var monthStr = String(currDate.getMonth() + 1);
	if(monthStr.length == 1) monthStr = "0"+monthStr;
	// Fill FILE_DATE with date variable
	metadata2DArray[metadataCol1DArray.indexOf("FILE_DATE")][1] = currDate;
	var path = "Kontrak"+"/"+CLIENT+"/"+yearStr+"-"+monthStr+"/"+EXT_REF_ID;
	logger.log("LOG INFO ===== path: "+path);
	
	// Ensure path exist
	var destFolder = ensureFolderExistByPath(path);
	var existingFile = destFolder.childByNamePath(filename);
	// Uploading document
	if(existingFile == null){
		logger.log("LOG ACTION ===== Existing document NOT found, begins uploading document");
		var upload = destFolder.createFile(filename) ;
		// upload.properties.content.write(content);
		// upload.properties.content.setEncoding("UTF-8");
		// upload.properties.content.guessMimetype(filename);
		upload.properties.content.write(content, false, true);

		// Change type and Inputting metadata
		upload.specializeType(typeQname);
		for(var i=0; i<metadata2DArray.length; i++){
			upload.properties[prefix+":"+metadata2DArray[i][0]] = metadata2DArray[i][1];
		}
		upload.save();
		result = "success";
		info = "Uploading successful with Content Type: "+typeQname;
	}else{
		logger.log("LOG ACTION ===== Existing document is found, begins updating document");
		existingFile.ensureVersioningEnabled(true, true);
		if (existingFile.isVersioned){
			var workingCopy = existingFile.checkout();
			// workingCopy.content = content;
			workingCopy.properties.content.write(content, false, true);
			workingCopy.checkin("", false); // false: minor changes, true: major changes
			
			// Inputting metadata
			if(existingFile.isSubType(typeQname) == false) existingFile.specializeType(typeQname);
			for(var i=0; i<metadata2DArray.length; i++){
				existingFile.properties[prefix+":"+metadata2DArray[i][0]] = metadata2DArray[i][1];
			}
			existingFile.save();
		}
		result = "success";
		info = "Updating existing document successful";
	}
}

var result = "nothing";
var info = "nothing";
logger.log("############ BEGINS UPLOAD WEBSCRIPT ############");
main();
model.result = result;
model.info = info;
logger.log("LOG INFO ===== result: " + result);
logger.log("LOG INFO ===== info: " + info);
logger.log("############ ENDS UPLOAD WEBSCRIPT ############\n\n");