
let name=/^[A-z ]{5,20}$/;

$("#txtName").on('keydown',function (event) {
    if(name.test($('#txtName').val())){
        $("#txtName").css('border','2px solid green');
    }else if(event.key=="Enter"){
        $('#txtAddress').focus();
    }
    else {
        $("#txtName").css('border','2px solid red');
        $('#txtName').focus();
    }
});

function textClear (){
        $("#txtNIC").val("");
        $("#txtName").val("");
        $("#txtAddress").val("");
        $("#txtSalary").val("");

    }
loadAllCustomers()
function loadAllCustomers() {
    let no=1;
    $('#tblCustomer').empty();
    $.ajax({
        method: "GET",
        crossDomain: true,
        url: "http://localhost:8080/api/v1/customer/all",
        success: function (res) {
            for (let i in res){
                let id = res[i].Id;
                let name = res[i].Name;
                let address = res[i].Address;
                let salary = res[i].Salary;
                let row=`<tr class="text-light"> <td>${no++}</td> <td>${id}</td><td>${name}</td><td>${address}</td><td>${salary.toFixed(2)}</td></tr>`;
                $('#tblCustomer').append(row);
            }
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    })

}
/*-----------------------Add Customer----------------------------------------------------------*/
//Add Customer
$("#btnAdd").click(function (event) {

    let name = $("#txtName").val();
    let address = $("#txtAddress").val();
    let salary = $("#txtSalary").val();
    event.preventDefault();
    if(name.trim().length>0&&address.trim().length>0&&salary.trim().length){
        // Get form
        let form = $('form').get(0);

        // Create an FormData object
        let data = new FormData(form);
        data.append("name",name);
        data.append("address",address);
        data.append("salary",salary);

        // disabled the submit button
        $("#btnImgUp").prop("disabled", true);

        $.ajax({
            type: "POST",
            enctype: 'multipart/form-data',
            url: "http://localhost:8080/api/v1/customer/add",
            data: data,
            processData: false,
            contentType: false,
            cache: false,
            timeout: 600000,
            success: function (data) {
                loadAllCustomers()
                console.log("SUCCESS : ", data);
                textClear()
            },
            error: function (e) {
                console.log("ERROR : ", e);
                /*       $("#btnImgUp").prop("disabled", false);*/

            }
        });
    }



});
/*--------------------------------------------------------------------*/

/*
$('#tblCustomer').on( 'click', 'tr', function () {
        console.log("Click Me")
        let id = $(this).children('td:eq(1)').text();
        let name = $(this).children('td:eq(2)').text();
        let address = $(this).children('td:eq(3)').text();
        let salary = $(this).children('td:eq(4)').text();

        $('#txtNIC').val(id);
        $('#txtName').val(name);
        $('#txtAddress').val(address);
        $('#txtSalary').val(salary);

        /!*  document.getElementById("btnDeleteCustomer").disabled = false;
          document.getElementById("btnUpdateCustomer").disabled = false;
          document.getElementById("btnSaveCustomer").disabled = true;*!/
    });*/

/*--------------------------------------------------------------------*/
function updateCustomer() {
    console.log("Delete")
    let id = $("#txtNIC").val();
    let name = $("#txtName").val();
    let address = $("#txtAddress").val();
    let salary = $("#txtSalary").val();
    console.log(id)
    console.log(name)
    console.log(address)
    console.log(salary)


    // Get form
    let form = $('form').get(0);


    let data = new FormData(form);
    data.append("Id", id);
    data.append("Name", name);
    data.append("Address", address);
    data.append("Salary", salary);


    $.ajax({
        type: "PUT",
        enctype: 'multipart/form-data',
        url: "http://localhost:8080/api/v1/customer/update",
        data: data,
        processData: false,
        contentType: false,
        cache: false,
        timeout: 600000,
        success: function (data) {
            loadAllCustomers()
            console.log("SUCCESS : ", data);
            textClear()
        },
        error: function (e) {
            console.log("ERROR : ", e);


        }
    });
}

/*-------------------------------------------------------------------------*/
//Delete Customer
$("#btnDelete").click(function (){
    let id = $("#txtNIC").val();
    // Get form
    let form = $('form').get(0);
    // Create an FormData object
    let data = new FormData(form);
    data.append("id",id);
    $.ajax({
        method:"DELETE",
        url:"http://localhost:8080/api/v1/customer/delete/"+id,
        data: data,
        processData: false,
        contentType: false,
        cache: false,
        timeout: 600000,
        success:function (res){
            alert("the customer is removed");
            loadAllCustomers()
            textClear()
        },
        error: function (ob, txtStatus, error) {
            console.log("===============");
            console.log(error);
            console.log("==================");
            console.log(txtStatus);
            console.log("================");
            console.log(ob.status);
        }
    });
});
/*----------Search Customer---------------------------*/
$("#btnSearch").click(function () {
    let id = $("#txtSearchID").val();
    console.log("Search Id IS "+id);
    $.ajax({
        method: "GET",
        url: "http://localhost:8080/api/v1/customer/search/"+id,
        success: function (res) {

            if(res.Name!=null){
                $("#txtNIC").val(res.Id);
                $("#txtName").val(res.Name);
                $("#txtAddress").val(res.Address);
                $("#txtSalary").val(res.Salary);
            }



        },
        error: function (ob, txtStatus, error) {
            // console.log(error);
            // console.log(txtStatus);
            // console.log(ob);
            alert("sorry!Customer Not found")
            textClear()
        }
    });
});
