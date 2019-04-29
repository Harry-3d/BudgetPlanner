function Msg() {
    alert('helloooooo')
};

function AddRow(name){
   var p = document.getElementById(name)
   /* var nameForm = document.createElement("input")
    nameForm.setAttribute('name', "static-name")
    nameForm.setAttribute('type', "text")
    */
    var br = document.createElement("br")

    var newRow = document.createElement('span')
    newRow.innerHTML = "<br> Name:<input type='text' name='" + name + "-name'> Cost: $ <input type='text' name='" + name +"-cost'>"

   // p.appendChild(br)
   // p.appendChild(nameForm)
    p.appendChild(newRow)

};

// document.write('FUUUUUUUUUUUUUUUUUUUCK')