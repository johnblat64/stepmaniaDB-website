<head>
    <link rel="stylesheet" type="text/css" href="/resources/p5.css">
</head>

<body>
    <h1 class="centered">www.StepmaniaDB.com</h1>
    <div class="centered">
        <small class="centered">Coming soon to a world wide web near you!</small>
        <small class = "centered"> This is a nice little site I made that can search through a database of stepmania song and chart metadata. Try out all the different form fields and enjoy using them. I especially like the stepstype search field. I hope you do too :) </small>
    </div>
    <div>
        <h2>{{.Title}}</h2>
        <ul>
            <li>Artist: {{.Artist}}</li>
            <li>Bpms: {{range $i, $a := .Bpms}} {{$a.Value}}, {{end}}</li>
            <li> TimeSignatures: 
                {{range $i, $a := .TimeSignatures}} 
                {{$a.Numerator}}/{{$a.Denominator}}
                {{end}}
            </li>

            
        </ul>
    </div>

    <h1 class="centered">You have successfully reached the end of the page</h1>
    
</body>