function traverse(nodes){
	for each(var node in nodes) {
		if (node.isContainer){
			// Node is a folder
			traverse(node.children);
		}else {
			// Node is a file
			allNodes.push(node);
		}
	}
}

// WEBSCRIPT for Getting Content Details (metadatas) and download link
var path = args.path; logger.log("LOG INFO ===== path: "+path);
var sortBy = args.sortBy; logger.log("LOG INFO ===== sortBy: "+sortBy);
var allNodes = new Array();
var folderNode = companyhome.childByNamePath(path);

traverse(folderNode.children);

model.result = "";
model.info = "";

if(allNodes[0] == null){
	result = "fail";
	info = "Get data failed due to incorrect Keyword";
	model.result = result; logger.log("LOG INFO ===== result: "+result);
	model.info = info; logger.log("LOG INFO ===== info: "+info);
}

// GET all results from search and push to an array for display later
var items = new Array();
for(var i=0; i<allNodes.length; i++){
	var node = allNodes[i];
	// logger.log("LOG INFO ===== node #"+i+": " + node.name);
	// logger.log("LOG INFO =============: " + node.isSubType("docpajak:spt"));
	items.push(node);
}

// Display all search results in model if there is any data
if(allNodes[0] != null){
	result = "success";
	info = "Get data success";
	model.result = result; logger.log("LOG INFO ===== result: "+result);
	model.info = info; logger.log("LOG INFO ===== info: "+info);

	model.data = items;
}

model.totalRecords = allNodes.length;