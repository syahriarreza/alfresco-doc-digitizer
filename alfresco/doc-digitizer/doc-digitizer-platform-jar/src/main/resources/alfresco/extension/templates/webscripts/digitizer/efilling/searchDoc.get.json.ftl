{
	"totalRecords": ${totalRecords},
	"items":
	[
		<#if data?exists>
			<#list data as item>
			{
				"name": "${item.name}",
				"nodeRef": "${item.nodeRef}",
				"title": "<#if item.properties["cm:title"]?exists>${item.properties["cm:title"]}</#if>",
				"description": "<#if item.properties["cm:description"]?exists>${item.properties["cm:description"]}</#if>",
				"type": "<#if item.properties["cm:type"]?exists>${item.properties["cm:type"]}</#if>",
				"created": "<#if item.properties["cm:created"]?exists>${xmldate(item.properties["cm:created"])}</#if>",
				"creator": "<#if item.properties["cm:creator"]?exists>${item.properties["cm:creator"]}</#if>",
				"modifiedOn": "<#if item.properties["cm:modified"]?exists>${xmldate(item.properties["cm:modified"])}</#if>",
				"modifiedByUser": "<#if item.properties["cm:modifier"]?exists>${item.properties["cm:modifier"]}</#if>",
				"accessed": "<#if item.properties["cm:accessed"]?exists>${item.properties["cm:accessed"]}</#if>",
				
				"NPWP": "<#if item.properties["docpajak:NPWP"]?exists>${item.properties["docpajak:NPWP"]}</#if>",
				"NAMA_WP": "<#if item.properties["docpajak:NAMA_WP"]?exists>${item.properties["docpajak:NAMA_WP"]}</#if>",
				"NM_JNS_SPT": "<#if item.properties["docpajak:NM_JNS_SPT"]?exists>${item.properties["docpajak:NM_JNS_SPT"]}</#if>",
				"LU_JNS_SPT": "<#if item.properties["docpajak:LU_JNS_SPT"]?exists>${item.properties["docpajak:LU_JNS_SPT"]}</#if>",
				"NM_KPP": "<#if item.properties["docpajak:NM_KPP"]?exists>${item.properties["docpajak:NM_KPP"]}</#if>",
				"THN_PJK": "<#if item.properties["docpajak:THN_PJK"]?exists>${item.properties["docpajak:THN_PJK"]}</#if>"
			}<#if item_has_next>,</#if>
			</#list>
		</#if>
	]
}