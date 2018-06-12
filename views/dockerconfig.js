var AuthCount = 1;
var HostCount = 1;
var ComponentCount = 1;

/**
* Function to add textbox element dynamically
* First we incrementing the counter and creating new div element with unique id
* Then adding new textbox element to this div and finally appending this div to main content.
*/
function addAuthData() {
    AuthCount++;
   var objNewDiv = document.createElement('div');
   objNewDiv.setAttribute('id', 'authdata_div_' + AuthCount);
   objNewDiv.innerHTML = ' <h3>AuthData'+AuthCount+'</h3>'+
   ' <label for="psw"><b>Password</b></label>'+
   '<input type="password" placeholder="Enter Password" class="password" >'+
   '<label for="username"><b>Username</b></label>'+
    '<input type="text" placeholder="Enter Username" class="username" >'+
    '<label for="auth"><b>Auth</b></label>'+
    '<input type="text" placeholder="Enter Auth" class="auth">'+
    '<label for="email"><b>Email</b></label>'+
    '<input type="text" placeholder="Enter Email" class="email" >';

   document.getElementById('authdata_container').appendChild(objNewDiv);
}

/**
* Function to remove textbox element dynamically
* check if counter is more than zero then remove the div element with the counter id and reduce the counter
* if counter is zero then show alert as no existing textboxes are there
*/
function removeAuthData() {
   if(0 < AuthCount) {
       document.getElementById('authdata_container').removeChild(document.getElementById('authdata_div_' + AuthCount));
       AuthCount--;
   } else {
       alert("No data to remove");
   }
}

function addHosts() {
    HostCount++;
    var objNewDiv = document.createElement('div');
   objNewDiv.setAttribute('id','hosts_div_'+HostCount);
   objNewDiv.innerHTML =' <h3>Host'+HostCount+'</h3>'+
   '<label for="protocol"><b>Protocol</b></label>'+
   '<input type="text" placeholder="Enter Protocol" class="protocol" >'+
   '<label for="apiversion"><b>Api Version</b></label>'+
   '<input type="text" placeholder="Enter Api Version" class="apiversion" >'+
   '<label for="hostype"><b>Host Type</b></label>'+
   '<input type="text" placeholder="Enter HostType" class="hostype" >'+
   '<label for="dockerversion"><b>Docker Version</b></label>'+
   '<input type="text" placeholder="Enter Docker Version" class="dockerversion" >'+
   '<label for="alias"><b>Alias</b></label>'+
   '<input type="text" placeholder="Enter Alias" class="alias" >'+
   '<label for="cert"><b>Cert Path For Docker Daemon</b></label>'+
   '<input type="text" placeholder="Enter Cert Path For Docker Daemon" class="cpath" >'+
   '<label for="ip"><b>Ip</b></label>'+
   '<input type="text" placeholder="Enter IP" class="Ip" >'+
   ' <label for="dockerport"><b>Docker Port</b></label>'+
   '<input type="text" placeholder="Enter Cert Path For Docker Daemon" class="dockerport" >';

document.getElementById('hosts_container').appendChild(objNewDiv);

}

function removeHosts() {
    if(0 < HostCount) {
        document.getElementById('hosts_container').removeChild(document.getElementById('hosts_div_' + HostCount));
        HostCount--;
    } else {
        alert("No data to remove");
    }
 }

function addComponent(){
    ComponentCount++;
    var objNewDiv = document.createElement('div');
   objNewDiv.setAttribute('id', 'component_div_' + ComponentCount);
   objNewDiv.classList.add('component_div')
   objNewDiv.innerHTML = '<h3>Component'+ ComponentCount +'</h3>' + '<label for="componentName"><b>Component name</b></label>'+
   '<input type="text" placeholder="Enter Component Name" class="componentName">'+'<p>Instances</p>'+
   '<div id="instance_div_component_'+ ComponentCount +'"> '+ '<div id="component_'+ ComponentCount +'_instance_1" class="instance container">'+
   '<h4>Component'+ ComponentCount +' Instance 1 </h4>'+
   '<label for="envMap"><b>envMap</b></label>'+
   '<input type="text" placeholder="Enter envMap" class="envmap">'+
   '<label for="portMapping"><b>portMapping</b></label>'+
   '<input type="text" placeholder="Enter portMapping" class="portmapping">'+
   '<label for="authId"><b>authId</b></label>'+
   '<input type="text" placeholder="Enter authId" class="authId">'+
   '<label for="hostId"><b>hostId</b></label>'+
   '<input type="text" placeholder="Enter hostId" class="hostId">'+
   '<label for="volumeMapping"><b>volumeMapping</b></label>'+
   '<input type="text" placeholder="Enter volumeMapping" class="volumemapping">'+
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
    objNewDiv.classList.add('instance')
    objNewDiv.classList.add('container')


    objNewDiv.innerHTML = '<h4>Component'+ ComponentNumber +'Instance'+ instanceCount +'</h4>'+
    '<label for="envMap"><b>envMap</b></label>'+
    '<input type="text" placeholder="Enter envMap" class="envmap">'+
    '<label for="portMapping"><b>portMapping</b></label>'+
    '<input type="text" placeholder="Enter portMapping" class="portmapping">'+
    '<label for="authId"><b>authId</b></label>'+
    '<input type="text" placeholder="Enter authId" class="authId">'+
    '<label for="hostId"><b>hostId</b></label>'+
    '<input type="text" placeholder="Enter hostId" class="hostId">'+
    '<label for="volumeMapping"><b>volumeMapping</b></label>'+
    '<input type="text" placeholder="Enter volumeMapping" class="volumemapping">'+
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
                var inst_envmap = $(value).find('.envmap')
                var inst_portmapping = $(value).find('.portmapping')
                var inst_authid = $(value).find('.auhtid')
                var inst_hostid = $(value).find('.hostid')
                var inst_volmapping = $(value).find('.volumemapping')
                var inst_volfrom = $(value).find('.volumesfrom')
                var inst_cexe = $(value).find('.commandtobeexecuted')
                var inst_links = $(value).find('.links')
                var inst_imgname = $(value).find('.imagename')
                var inst_tag = $(value).find('.tag')
                var inst_hostsmapping = $(value).find('.hostsmapping')
                var inst_name = $(value).find('.name')


                
                
                var ienvmap = $(inst_envmap[0]).val();
                var iportmapping = $(inst_portmapping[0]).val();
                var iauthid = $(inst_authid[0]).val();
                var ihostid = $(inst_hostid[0]).val();
                var ivolmapping = $(inst_volmapping[0]).val();
                var ivolfrom = $(inst_volfrom[0]).val();

                var icexe = $(inst_cexe[0]).val();
                var ilinks = $(inst_links[0]).val();
                var i_imgname = $(inst_imgname[0]).val();
                var itag = $(inst_tag[0]).val();
                var ihostsmapping = $(inst_hostsmapping[0]).val();
                var iname = $(inst_name[0]).val();

            var instance = {}

            instance.envMap = ienvmap;
            instance.portMapping = iportmapping;
            instance.authId = iauthid;
            instance.hostId = ihostid;
            instance.volumeMapping = ivolmapping;
            instance.volumesFrom = ivolfrom;
            instance.

            instances.push(instance)
        })    

        component.instances = instances
        components.push(component) 
    })
    jsonObject.components = components;
            

            jQuery.ajax({
                url: 'http://localhost:8080/docker/config',
                type: "POST",
                data:JSON.stringify(jsonObject),
                dataType: "jsonp",
                success: function(success) {
                console.log(success)
                }
      });
    })
})