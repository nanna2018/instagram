
$(document).ready(function() {
    console.log("hola")
    ActualizarFotos();
    $("#txtEmail").keyup(function(event) {
        if (event.keyCode === 13) {
             $("btnEnviar").click();
        }
    });
//Registro
    $("#btnEnviar").click(function() {
         var name = $("#txtTexto").val()
         var password = $("#txtPassword").val()
         var email = $("#txtEmail").val()
 
        console.log(name, password,email );
 
        var envio = {
             name: name,
             password: password,
             email: email
        };
 
         $.post({
             url:"/envio",
             data: JSON.stringify(envio),
             success: function(data, status, jqXHR) {
                 console.log(data);
                 $("#txtTexto").val('')
                 $("#txtPassword").val('')
                 $("#txtEmail").val('')
             },
             dataType: "json"
 
         }).done(function(data) {
             console.log("Petición realizada");
             //ActualizarHistorial();
         
         }).fail(function(data) {
             console.log("Petición fallida");
         
         }).always(function(data){
             console.log("Petición completa");
         });
    });
//Login
    $("#btnLogin").click(function() {
        var name = $("#txtTexto").val()
        var password = $("#txtPassword").val()
       console.log(name, password );

       var login = {
            name: name,
            password: password,
       };

        $.post({
            url:"/login",
            data: JSON.stringify(login),
            method:"POST",
            success: function(data, status, jqXHR) {
                console.log(data);
                
            },
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");
            if(data==true){
               window.location.href="/perfil";
            }
           
            //ActualizarHistorial();
        
        }).fail(function(data) {
            console.log("Petición fallida");
            
        
        }).always(function(data){
            console.log("Petición completa");
        });
   });


   //fotos
  
   
   //Ajax para mostrar las fotos
function ActualizarFotos() {   
    $.ajax({
        url: "/lista",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
            Historial_Fotos(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}

function Historial_Fotos(array) {
    
    var div = $("#fotos");
    div.children().remove();
    if(array != null && array.length > 0) {
      

        for(var x = 0; x < array.length; x++) {
            div.append( 
            "<div>"
                +"<img src='/files/"+array[x].ID+"' width='300px' height='180px'>"+
                "<p>"+array[x].Name+"</p>"+
            "</div>");
        }
    } else {
        div.append('<div colspan="3">No hay registros de hoy</div>');
        
    }
}




})