{{define "content"}}

    <div>
        <h1>{{.Title}}</h1>
        <ul>
            <li>Pack: <a href="/packs/{{.PackId}}">{{.PackName}}</a></li>
            <li>Artist: {{.Artist}}</li>
            <li>Bpms: {{range $i, $a := .Bpms}} {{$a.Value}}, {{end}}</li>
            <li> TimeSignatures: 
                {{range $i, $a := .TimeSignatures}} 
                {{$a.Numerator}}/{{$a.Denominator}}
                {{end}}
            </li>

            <div>
                <h2> Charts </h2>
                    {{range $i, $chart := .Charts}}
                    <h3> {{$chart.Description}} </h3>
                        <ul>
                            <li>Steps Type: {{$chart.StepsType}}</li>
                            <li>Number of Stops: {{$chart.StopsCount}}</li>
                            <li> Number of Delays: {{$chart.DelaysCount}}</li>
                            <li> Number of Warps: {{$chart.WarpsCount}}</li>
                            <li> Number of Scrolls: {{$chart.ScrollsCount}}</li>
                            <li> Number of Fakes: {{$chart.FakesCount}}</li>
                            <li> Number of Speeds: {{$chart.SpeedsCount}}</li>
                            <li> Meter: {{$chart.Meter}} </li>
                        </ul>
                    {{end}}
            </div>

            
        </ul>
    </div>
    
{{end}}