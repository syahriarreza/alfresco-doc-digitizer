// WEBSCRIPT for Getting Content Details (metadatas) and download link
var keyword = args.keyword; logger.log("LOG INFO ===== keyword: "+keyword);
var sortBy = args.sortBy; logger.log("LOG INFO ===== sortBy: "+sortBy);
var allNodes;

if(sortBy == null || sortBy == "") allNodes = search.luceneSearch(keyword);
else allNodes = search.luceneSearch(keyword, sortBy, true);

model.result = "";
model.info = "";

if(allNodes[0] == null){
	result = "fail";
	info = "Get data failed due to incorrect Keyword";
}

// Display all search results in model
if(allNodes[0] != null){
	result = "success";
	info = "Get data success";

	model.data = allNodes;
}

model.totalRecords = allNodes.length;
model.result = result; logger.log("LOG INFO ===== result: "+result);
model.info = info; logger.log("LOG INFO ===== info: "+info);