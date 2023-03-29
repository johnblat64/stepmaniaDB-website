<body>
    <h1 class="centered">www.StepmaniaDB.com</h1>
    <div class="boxed">
        <h1>{{.Title}}</h1>
        <ul>
            <li>Subtitle: {{.Subtitle}}</li>
            <li>Artist: {{.Artist}}</li>
            <li>Bpms: {{range $i, $a := .Bpms}} {{$a}}, {{end}}</li>
            <li> TimeSignatures: 
                {{range= $i, $a := .TimeSignatures}} 
                {{$a.Numerator}}/{{$a.Denominator}}
                {{end}}
            </li>

            
        </ul>
    </div>
    
</body>