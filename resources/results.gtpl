<head>
<link rel="stylesheet" type="text/css" href="/resources/style.css">
</head>


<body class="bg">

<h1 class="centered">www.StepmaniaDB<a style="text-decoration: none; color: green" href="https://youtu.be/ruSjI7r1sO0?t=294">.</a>com</h1>
<div class="centered">
    <small class="centered">Coming soon to a world wide web near you!</small>
    <br>
    <img class="centered" src="/resources/ddr.gif" alt="ddr"> 

</div>
<h2>Search Results</h2> 

    {{range .SongPage.Songs}}
    <a style="text-decoration: none; color: black" href="song">
        <span style="display: block; " class="boxed" >
            <ul >
            
                <h3>{{.Title}} </h3>
                <li>Artist: {{.Artist}}</li>
                <li> Bpms: {{range $i, $a := .Bpms}} {{$a}}, {{end}}</li>
            </ul>
        </span>
    </a>
   
    {{end}}

    <h1 class="centered">You have successfully reached the end of the page</h1>


</body>
