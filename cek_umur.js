<script>

var umur = [20, 30, 40, 45];

//
function cekumur(age) {
    return age >= 22;
}

//
function myFunction() {
    document.getElementById("demo").innerHTML = umur.filter(cekumur);
}<br /></script>