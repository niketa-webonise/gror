<!DOCTYPE html>
<html>
  <head>
      <meta charset="utf-8">
      
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css">
      <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
      
<style>
p{
  color:red;
}

body {font-family: Arial, Helvetica, sans-serif;}
* {box-sizing: border-box}
/* Full-width input fields */
input[type=text], input[type=password] {
    width: 100%;
    padding: 15px;
    margin: 5px 0 22px 0;
    display: inline-block;
    border: none;
    background: #f1f1f1;
}

/* Add a background color when the inputs get focus */
input[type=text]:focus, input[type=password]:focus {
    background-color: #ddd;
    outline: none;
}

/* Set a style for all buttons */
button {
    background-color: #4CAF50;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    cursor: pointer;
    width: 100%;
    opacity: 0.9;
}

button:hover {
    opacity:1;
}


/* Extra styles for the cancel button */
.cancelbtn {
    padding: 14px 20px;
    background-color: #f44336;
}

/* Float cancel and signup buttons and add an equal width */
.cancelbtn, .signupbtn {
  float: left;
  width: 50%;
}

/* Add padding to container elements */
.container {
    padding: 16px;
}

/* The Modal (background) */
.modal {
    display: none; /* Hidden by default */
    position: fixed; /* Stay in place */
    z-index: 1; /* Sit on top */
    left: 0;
    top: 0;
    width: 100%; /* Full width */
    height: 100%; /* Full height */
    overflow: auto; /* Enable scroll if needed */
    background-color: #474e5d;
    padding-top: 50px;
}

/* Modal Content/Box */
.modal-content {
    background-color: #fefefe;
    margin: 5% auto 15% auto; /* 5% from the top, 15% from the bottom and centered */
    border: 1px solid #888;
    width: 80%; /* Could be more or less, depending on screen size */
}

/* Style the horizontal ruler */
hr {
    border: 1px solid #f1f1f1;
    margin-bottom: 25px;
}
 
/* The Close Button (x) */
.close {
    position: absolute;
    right: 35px;
    top: 15px;
    font-size: 40px;
    font-weight: bold;
    color: #f1f1f1;
}

.close:hover,
.close:focus {
    color: #f44336;
    cursor: pointer;
}

/* Clear floats */
.clearfix::after {
    content: "";
    clear: both;
    display: table;
}

/* Change styles for cancel button and signup button on extra small screens */
@media screen and (max-width: 300px) {
    .cancelbtn, .signupbtn {
       width: 100%;
    }
}
</style>
<script type="text/javascript">

    var intAuth = 1;
    var intHost = 1;
    var ComponentCount = 1;


function addAuth() {

    intAuth++;
    var objNewDiv = document.createElement('div');
    objNewDiv.setAttribute('id', 'divauth_' + intAuth);
    objNewDiv.innerHTML = '<h3>AuthData ' + intAuth + '</h3>'+
    '<label><b>Password</b>'+
    '</label><input type="password" placeholder="Enter Password" name="password" class="pswtb">'+
    '<label><b>Key</b>'+
    '</label><input type="password" placeholder="Enter Password" class="keytb" name="key">'+
    '<label><b>Username</b></label>'+
    '<input type="text" placeholder="Enter Username" class="unametb" name="username">'+
    '<label><b>Auth</b></label>'+
    '<input type="text" placeholder="Enter Auth" class="authtb" name="auth" >'+
    '<label><b>Email</b></label>'+
    '<input type="text" placeholder="Enter Email" class="emailtb" name="email">';
    document.getElementById('authdata_container').appendChild(objNewDiv);
}

function removeAuth() {
    if(0 < intAuth) {
        document.getElementById('authdata_container').removeChild(document.getElementById('divauth_' + intAuth));
        intAuth--;
    } else {
        alert("No data to remove");
    }
}

function addHost() {
  intHost++;
    var objNewDiv = document.createElement('div');
    objNewDiv.setAttribute('id', 'divhost_' + intHost);
    objNewDiv.innerHTML = '<h3>Hosts ' + intHost + '</h3>'+
    '<label><b>Protocol</b></label><input type="text" placeholder="Enter Protocol" class="protb" name="protocol">'+
    '<label><b>Api Version</b></label>'+
    '<input type="text" placeholder="Enter Api Version" class="Apivertb" name="apiVersion">'+
    '<label><b>Host Type</b></label>'+
    '<input type="text" placeholder="Enter Host Type" class="hostypetb" name="hostType" >'+
    '<label><b>Docker Version</b></label>'+
    '<input type="text" placeholder="Enter Docker Version" class="dockvertb" name="dockerVersion">'+
    '<label><b>Alias</b></label>'+
    '<input type="text" placeholder="Enter Alias" class="aliastb" name="alias">'+
    '<label><b>Cert Path For Docker Daemon</b></label>'+
    '<input type="text" placeholder="Enter Cert Path For Docker Daemon" class="cpathtb" name="certPathForDockerDaemon">'+
    '<label><b>Ip</b></label>'+
    '<input type="text" placeholder="Enter Ip" class="Iptb" name="ip">'+
    '<label><b>Docker Port</b></label>'+
    '<input type="text" placeholder="Enter Docker Port" class="DocPorttb" name="dockerPort">';
    document.getElementById('hostdata_container').appendChild(objNewDiv);
}

function removeHost() {
    if(0 < intHost) {
        document.getElementById('hostdata_container').removeChild(document.getElementById('divhost_' + intHost));
        intHost--;
    } else {
        alert("No data to remove");
    }
}

function addComponent(){
    ComponentCount++;
    var objNewDiv = document.createElement('div');
   objNewDiv.setAttribute('id', 'component_div_' + ComponentCount);
   objNewDiv.classList.add('component_div');
   objNewDiv.innerHTML = '<h3>Component'+ ComponentCount +'</h3>' + '<label for="componentName"><b>Component name</b></label>'+
   '<input type="text" placeholder="Enter Component Name" class="componentName">'+'<p>Instances</p>'+
   '<div id="instance_div_component_'+ ComponentCount +'"> '+ '<div id="component_'+ ComponentCount +'_instance_1" class="instance container">'+
   '<h4>Component'+ ComponentCount +' Instance 1 </h4>'+
   '<label for="envMap"><b>envMap</b></label>'+
   '<input type="text" placeholder="Enter CASSANDRA_BROADCAST_ADDRESS" class="cass">'+
    '<input type="text" placeholder="Enter CASSANDRA_SEEDS" class="cass_seeds">'+
   '<label for="portMapping"><b>portMapping</b></label>'+
   '<input type="text" placeholder="Enter portMapping" class="portmapping">'+
   '<label for="authId"><b>authId</b></label>'+
   '<input type="text" placeholder="Enter authId" class="authId">'+
   '<label for="hostId"><b>hostId</b></label>'+
   '<input type="text" placeholder="Enter hostId" class="hostId">'+
   '<label for="volumeMapping"><b>volumeMapping</b></label>'+
   '<input type="text" placeholder="Enter /home/ubuntu/cass-data" class="cass-data">'+
   '<input type="text" placeholder="Enter /home/ubuntu/cass-config" class="cass-config">'+
   '<label for="volumesFrom"><b>volumesFrom</b></label>'+
    '<input type="text" placeholder="Enter volumesFrom" class="volumesfrom">'+
    '<label for="commandToBeExecuted"><b>commandToBeExecuted</b></label>'+
    '<input type="text" placeholder="Enter commandToBeExecuted" class="commandtobeexecuted">'+
    '<label for="links"><b>links</b></label>'+
    '<input type="text" placeholder="Enter links" class="links">'+
    '<label for="imageName"><b>imageName</b></label>'+
    '<input type="text" placeholder="Enter imageName" class="imagename">'+
    '<label for="tag"><b>tag</b></label>'+
    '<input type="text" placeholder="Enter tag" class="tag">'+
    '<label for="hostsMapping"><b>hostsMapping</b></label>'+
    '<input type="text" placeholder="Enter hostsmapping" class="hostsmapping">'+
    '<label for="instanceName"><b>name</b></label>'+
    '<input type="text" placeholder="Enter name" class="name">'+
    '<p>'+
        '<a href="javascript:void(0);" onclick="addInstance('+ ComponentCount +');">Add Instance</a>'+
            '<a href="javascript:void(0);" onclick="removeInstance('+ ComponentCount +');">Remove Instance</a>'+
    '</p>'+
'</div>'+
'</div>';
    
    document.getElementById('component_container').appendChild(objNewDiv);
}

function removeComponent() {
    if(0 < ComponentCount) {
        document.getElementById('component_container').removeChild(document.getElementById('component_div_' + ComponentCount));
        ComponentCount--;
    } else {
        alert("No data to remove");
    }
 }


 function addInstance(ComponentNumber){

    var instanceContainer = document.getElementById("instance_div_component_"+ComponentNumber);

    var children = $('#instance_div_component_'+ ComponentNumber).children('.instance');

    var instanceCount = children.length;
    instanceCount++

    var objNewDiv = document.createElement('div');
    objNewDiv.setAttribute('id','component_'+ ComponentNumber +'_instance_'+ instanceCount );
    objNewDiv.classList.add('instance');
    objNewDiv.classList.add('container');


    objNewDiv.innerHTML = '<h4>Component'+ ComponentNumber +'Instance'+ instanceCount +'</h4>'+
    '<label for="envMap"><b>envMap</b></label>'+
    '<input type="text" placeholder="Enter CASSANDRA_BROADCAST_ADDRESS" class="cass">'+
    '<input type="text" placeholder="Enter CASSANDRA_SEEDS" class="cass_seeds">'+
    '<label for="portMapping"><b>portMapping</b></label>'+
    '<input type="text" placeholder="Enter portMapping" class="portmapping">'+
    '<label for="authId"><b>authId</b></label>'+
    '<input type="text" placeholder="Enter authId" class="authId">'+
    '<label for="hostId"><b>hostId</b></label>'+
    '<input type="text" placeholder="Enter hostId" class="hostId">'+
    '<label for="volumeMapping"><b>volumeMapping</b></label>'+
    '<input type="text" placeholder="Enter /home/ubuntu/cass-data" class="cass-data">'+
   '<input type="text" placeholder="Enter /home/ubuntu/cass-config" class="cass-config">'+
    '<label for="volumesFrom"><b>volumesFrom</b></label>'+
    '<input type="text" placeholder="Enter volumesFrom" class="volumesfrom">'+
    '<label for="commandToBeExecuted"><b>commandToBeExecuted</b></label>'+
    '<input type="text" placeholder="Enter commandToBeExecuted" class="commandtobeexecuted">'+
    '<label for="links"><b>links</b></label>'+
    '<input type="text" placeholder="Enter links" class="links">'+
    '<label for="imageName"><b>imageName</b></label>'+
    '<input type="text" placeholder="Enter imageName" class="imagename">'+
    '<label for="tag"><b>tag</b></label>'+
    '<input type="text" placeholder="Enter tag" class="tag">'+
    '<label for="hostsMapping"><b>hostsMapping</b></label>'+
    '<input type="text" placeholder="Enter hostsmapping" class="hostsmapping">'+
    '<label for="instanceName"><b>name</b></label>'+
    '<input type="text" placeholder="Enter name" class="name">'+
    '<p>'+
        '<a href="javascript:void(0);" onclick="addInstance('+ ComponentNumber +');">Add Instance</a>'+
            '<a href="javascript:void(0);" onclick="removeInstance('+ ComponentNumber +');">Remove Instance</a>'+
    '</p>'+
'</div>'+
'</div>';

    document.getElementById('instance_div_component_'+ ComponentNumber).appendChild(objNewDiv);
 }

function removeInstance(ComponentNumber) {
    var instanceContainer = document.getElementById('instance_div_component_'+ComponentNumber);
    var instanceCount = 0;
    var children = $('#instance_div_component_'+ComponentNumber).children('.instance');
    instanceCount = children.length
    instanceContainer.removeChild(document.getElementById('component_'+ ComponentNumber +'_instance_'+ instanceCount));
}

$(document).ready(function(){
    $('#gror_form').on('submit', function(e) {
         e.preventDefault();
        
         var jsonObject = {}
         var components = []; 
         var auths=[];
         var hosts=[];

          var systemInfo = {}
        
        systemInfo.grorVersion=$("#gversion").val();
        systemInfo.name=$("#gverName").val();
        jsonObject.systemInfo =systemInfo ;

         var authChildren = $('#authdata_container').children()
        $.each(authChildren, function( index, value ) {
            var auth = {}
            auth.password=$($(value).find('.pswtb')[0]).val();
            auth.key=$($(value).find('.keytb')[0]).val();
            auth.username=$($(value).find('.unametb')[0]).val();
            auth.auth=$($(value).find('.authtb')[0]).val();
            auth.email=$($(value).find('.emailtb')[0]).val();
            auths.push(auth)
        })
        jsonObject.authData = auths;


        var hostChildren = $('#hostdata_container').children()
        $.each(hostChildren, function( index, value ) {
            var host={}
            host.protocol=$($(value).find('.protb')[0]).val();
            host.apiVersion=$($(value).find('.Apivertb')[0]).val();
            host.hostType=$($(value).find('.hostypetb')[0]).val();
            host.dockerVersion=$($(value).find('.dockvertb')[0]).val();
            host.alias=$($(value).find('.aliastb')[0]).val();
            host.certPathForDockerDaemon=$($(value).find('.cpathtb')[0]).val();
            host.ip=$($(value).find('.Iptb')[0]).val();
            host.dockerPort=$($(value).find('.DocPorttb')[0]).val();
            hosts.push(host)
        })
        jsonObject.hosts = hosts;

         
        var componentChildren = $('#component_container').children()
      
           $.each(componentChildren, function( index, value ) {

                var componentNo = index + 1;
                var cmp_name_input = $(value).find('.componentName')
                var cname = $(cmp_name_input[0]).val();

                var component = {}
                component.name = cname;
                var instances = [];

                var instanceChildren = $('#instance_div_component_'+componentNo).children()
            
                $.each(instanceChildren, function(index,value){
                var instanceNo = index + 1;
                var inst_env_cass = $(value).find('.cass')
                var inst_env_cass_seeds = $(value).find('.cass_seeds')

                var inst_portmapping = $(value).find('.portmapping')
                var inst_authid = $(value).find('.authid')
                var inst_hostid = $(value).find('.hostid')

                var inst_volmapping_cassdata = $(value).find('.cass-data')
                var inst_volmapping_cassconfig = $(value).find('.cass-config')

                var inst_volfrom = $(value).find('.volumesfrom')
                var inst_cexe = $(value).find('.commandtobeexecuted')
                var inst_links = $(value).find('.links')
                var inst_imgname = $(value).find('.imagename')
                var inst_tag = $(value).find('.tag')
                var inst_hostsmapping = $(value).find('.hostsmapping')
                var inst_name = $(value).find('.name')

  
                var ienvmapcass = $(inst_env_cass[0]).val();
                var ienvmapcass_seeds = $(inst_env_cass_seeds[0]).val();

                var iportmapping = $(inst_portmapping[0]).val();
                var iauthid = $(inst_authid[0]).val();
                var ihostid = $(inst_hostid[0]).val();

                var ivolmapping_cassdata = $(inst_volmapping_cassdata[0]).val();
                var ivolmapping_cassconfig = $(inst_volmapping_cassconfig[0]).val();
                
                var ivolfrom = $(inst_volfrom[0]).val();
                var icexe = $(inst_cexe[0]).val();
                var ilinks = $(inst_links[0]).val();
                var i_imgname = $(inst_imgname[0]).val();
                var itag = $(inst_tag[0]).val();
                var ihostsmapping = $(inst_hostsmapping[0]).val();
                var iname = $(inst_name[0]).val();

            var instance = {}
            var envmap = {}
            var volumemapping = {}
                  
            envmap.CASSANDRA_BROADCAST_ADDRESS = ienvmapcass;
            envmap.CASSANDRA_SEEDS = ienvmapcass_seeds;

            instance.envMap = envmap;
            instance.portMapping = iportmapping;
            instance.authId = iauthid;
            instance.hostId = ihostid;

            volumemapping["/home/ubuntu/cass-data"] = ivolmapping_cassdata;
            volumemapping["/home/ubuntu/cass-config"] = ivolmapping_cassconfig;
            
            instance.volumeMapping = volumemapping;
            instance.volumesFrom = ivolfrom;
            instance.commandToBeExecuted = icexe;
            instance.links = ilinks;
            instance.imageName = i_imgname;
            instance.tag = itag;
            instance.hostsMapping = ihostsmapping;
            instance.name = iname;

            instances.push(instance)
        })    

        component.instances = instances
        components.push(component) 
    })
    jsonObject.components = components;
            

            $.ajax({
                url: 'http://localhost:8080/docker/config',
                contentType : "application/json",
                type: "POST",
                dataType: "json",
                data:JSON.stringify(jsonObject),
                success: function(result) {
                alert(result.message)
                window.location.href="/docker/config"
                },
                error: function(xhr,type,exception){
                    alert("ajax error response type"+type)
                }
      });
    })
})
</script>
</head>
<body>
<form  id="gror_form"  style="border:1px solid #ccc">
  <div class="container">
        <h1>Docker</h1>
        <h3>SystemInfo</h3>
        <hr>
            <div id="SystemInfo_div" >
                <label for="psw-repeat">
                    <b>grorVersion</b>
                </label>
                <input type="text" placeholder="Enter grorversion" name="grorversion" id="gversion" >
                
                <label for="psw-repeat">
                    <b>Name</b>
                </label>
                <input type="text" placeholder="Enter version Name" name="versionname" id="gverName">
            </div>
    

    
    <div id="authdata_container" class="auth_div">
        <div id="divauth_1">
            <h3>AuthData 1</h3>  
            <label>
                <b>Password</b>
            </label>
            <input type="password" placeholder="Enter Password" name="password"  class="pswtb">
           
            <label>
                <b>Key</b>
            </label>
            <input type="text" placeholder="Enter Key" name="key" class="keytb" >
            
            <label>
                <b>Username</b>
            </label>
            <input type="text" placeholder="Enter Username" name="username" class="unametb" >
            
            <label>
                <b>Auth</b>
            </label>
            <input type="text" placeholder="Enter Auth" name="auth" class="authtb" >
            
            <label>
                <b>Email</b>
            </label>
            <input type="text" placeholder="Enter Email" name="email"  class="emailtb">
            
       
        </div>
    </div>

        <p>
            <a href="javascript:void(0);" onclick="addAuth();">Add</a>
            <a href="javascript:void(0);" onclick="removeAuth();">Remove</a>
        </p>


    <div id="hostdata_container">
        <div id="divhost_1">
    
            <h3>Hosts 1</h3>
    
            <label>
                <b>Protocol</b>
            </label>
            <input type="text" placeholder="Enter Protocol" name="protocol" class="protb" >
    
            <label>
                <b>Api Version</b>
            </label>
            <input type="text" placeholder="Enter Api Version" name="apiVersion" class="Apivertb" >
    
            <label>
                <b>Host Type</b>
            </label>
            <input type="text" placeholder="Enter Host Type" name="hostType" class="hostypetb">
            <label>
                <b>Docker Version</b>
            </label>
            <input type="text" placeholder="Enter Docker Version" name="dockerVersion" class="dockvertb">
            
            <label>
                <b>Alias</b>
            </label>
            <input type="text" placeholder="Enter Alias" name="alias" class="aliastb">
    
            <label>
            <b>Cert Path For Docker Daemon</b>
            </label>
            <input type="text" placeholder="Enter Cert Path For Docker Daemon" name="certPathForDockerDaemon" class="cpathtb">
    
            <label>
                <b>Ip</b>
            </label>
            <input type="text" placeholder="Enter Ip" name="ip" class="Iptb" >
    
            <label>
                <b>Docker Port</b>
            </label>
            <input type="text" placeholder="Enter Docker Port" name="dockerPort" class="DocPorttb" >

        
        </div>
    </div>
    <p>
        <a href="javascript:void(0);" onclick="addHost();">Add</a>
        <a href="javascript:void(0);" onclick="removeHost();">Remove</a>
    </p>
    
    <p>components</p>
    <div id="component_container">
        <div id="component_div_1" class="component_div">
            <h3>Component 1</h3>
            <label for="componentName">
                <b>Component name</b>
            </label>
            <input type="text" placeholder="Enter Component Name" class="componentName">

                <p>instances</p>
                <div id="instance_div_component_1">
                    
                    <div id="component_1_instance_1" class="instance container">
                        <label for="envMap">
                            <b>envMap</b>
                        </label>
                        <input type="text" placeholder="Enter CASSANDRA_BROADCAST_ADDRESS" class="cass">

                        <label for="envMap">
                            <b>envMap</b>
                        </label>
                        <input type="text" placeholder="Enter CASSANDRA_SEEDS" class="cass_seeds">

                        <label for="portMapping">
                            <b>portMapping</b>
                        </label>
                        <input type="text" placeholder="Enter portMapping" class="portmapping">

                        <label for="authId">
                            <b>authId</b>
                        </label>
                        <input type="text" placeholder="Enter authId" class="authid">

                        <label for="hostId">
                            <b>hostId</b>
                        </label>
                        <input type="text" placeholder="Enter hostId" class="hostid">

                        <label for="volumeMapping">
                            <b>volumeMapping</b>
                        </label>
                        <input type="text" placeholder="Enter /home/ubuntu/cass-data" class="cass-data">
                        <input type="text" placeholder="Enter /home/ubuntu/cass-config" class="cass-config">

                        <label for="volumesFrom">
                            <b>volumesFrom</b>
                        </label>
                        <input type="text" placeholder="Enter volumesFrom" class="volumesfrom">

                        <label for="commandToBeExecuted">
                            <b>commandToBeExecuted</b>
                        </label>
                        <input type="text" placeholder="Enter commandToBeExecuted" class="commandtobeexecuted">

                        <label for="links">
                            <b>links</b>
                        </label>
                        <input type="text" placeholder="Enter links" class="links">

                        <label for="imageName">
                            <b>imageName</b>
                        </label>
                        <input type="text" placeholder="Enter imageName" class="imagename">

                        <label for="tag">
                            <b>tag</b>
                        </label>
                        <input type="text" placeholder="Enter tag" class="tag">

                        <label for="hostsMapping">
                            <b>hostsMapping</b>
                        </label>
                        <input type="text" placeholder="Enter hostsmapping" class="hostsmapping">

                        <label for="instanceName">
                            <b>name</b>
                        </label>
                        <input type="text" placeholder="Enter name" class="name">

                        <p>
                            <a href="javascript:void(0);" onclick="addInstance(1);">Add Instance</a>
                        <a href="javascript:void(0);" onclick="removeInstance(1);">Remove Instance</a>
                        </p>
                    </div>
                </div>
        </div>
    </div>
    <p>
        <a href="javascript:void(0);" onclick="addComponent();">Add Component</a>
       <a href="javascript:void(0);" onclick="removeComponent();">Remove Component</a>
     </p>

    <div class="clearfix">
      <button type="submit" class="signupbtn">Submit</button>
    </div>

  </div>
</form>
</body>
</html>