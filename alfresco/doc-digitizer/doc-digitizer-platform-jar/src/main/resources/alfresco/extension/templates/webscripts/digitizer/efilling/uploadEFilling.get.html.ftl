<html>
 <head> 
   <title>Upload Web Script E-Filling</title> 
   <link rel="stylesheet" href="${url.context}/css/main.css" TYPE="text/css">
 </head>
 <body>
   <table>
     <tr>
       <td><img src="${url.context}/images/logo/AlfrescoLogo32.png" alt="Alfresco" /></td>
       <td><nobr>Upload Web Script E-Filling</nobr></td>
     </tr>
     <tr><td><td>Alfresco ${server.edition} v${server.version}
   </table>
   <p>
   <table>
     <form action="${url.service}" method="post" enctype="multipart/form-data" accept-charset="utf-8">
       <tr><td>File:</td><td><input type="file" name="file"></td></tr>
	   <tr><td>Metadata:</td><td><input name="metadata"></td></tr>
       <tr><td></td></tr>
       <tr><td><input type="submit" name="submit" value="Upload"></td></tr>
     </form>
   </table>
   <br/>
   example: FILE_TYPE:-SPT Tahunan,DOC_TYPE:-Formulir,CLIENT:-PROSIA,EXT_REF_ID:-1234567,FILE_DATE:-2015/01/28
 </body>
</html>