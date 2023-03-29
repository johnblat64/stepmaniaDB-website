<style>
body {
    background-color: aquamarine;
}
.boxed {
    border: 2px solid green ;
    margin-bottom: 5px;
}

.boxed:hover {
    background-color: green;
    color: pink
}

.centered {
    text-align: center;
    align-self: center;
}

h1 {
    color: green
}
</style>


<head>
    <link rel="stylesheet" href="style.css">
</head>


<body>

<h1 class="centered">www.StepmaniaDB.com</h1>
<div class="centered">
    <small class="centered">Coming soon to a world wide web near you!</small>
    <br>
    <img class="centered" src="/resources/ddr.gif" alt="ddr"> 

</div>
<h2>Search Results</h2> 

    {{range .Songs}}
    <div class="boxed">
        <ul >
            <h3>{{.Title}} </h3>
            <li>Artist: {{.Artist}}</li>
            <li> Bpms: {{range $i, $a := .Bpms}} {{$a}}, {{end}}</li>
        </ul>

    </div>
    {{end}}

</body>
