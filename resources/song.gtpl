<head>
    <link rel="stylesheet" type="text/css" href="/resources/style.css">
</head>

<body>
    <h1 class="centered">www.StepmaniaDB.com</h1>
    <div class="centered">
        <small class="centered">Coming soon to a world wide web near you!</small>
    </div>
    <div>
        <h2>{{.Title}}</h2>
        <ul>
            <li>Artist: {{.Artist}}</li>
            <li>Bpms: {{range $i, $a := .Bpms}} {{$a}}, {{end}}</li>
            <li> TimeSignatures: 
                {{range $i, $a := .TimeSignatures}} 
                {{$a.Numerator}}/{{$a.Denominator}}
                {{end}}
            </li>

            
        </ul>
    </div>

    <h1 class="centered">You have successfully reached the end of the page</h1>
    
</body>