function main(){
	// Getting Data
	var filePath = args.filePath;
	
	// Checking File Existence
	var fileNode = companyhome.childByNamePath(filePath);
	if(fileNode == null){ result = "NO"; info = "File is NOT found with path: "+filePath; }
	else{ result = "YES"; info = "File exist in \""+filePath+"\""; }
}

var result = "nothing";
var info = "nothing";
main();
model.result = result;
// model.info = info;
logger.log("LOG INFO = result: "+result+". "+info);