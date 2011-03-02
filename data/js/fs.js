
var ListDir = function(dir) {
	$.ajax({
		type: "GET",
		url: "ls/"+dir,
		success: ListSuccess
	});
}

var ListSuccess = function(data, textStatus, jqXHR) {
	$("#filebrowser").treeTable("destroy")
	$("#filebrowser").html(data)
	$("#filebrowser").treeTable()	

	aceEditor.getSession().setValue(data)
}

$(document).ready(function(){
	ListDir(".")
});