{{define "content"}}

    <div>
        <h1>{{.Title}}</h1>
        <img src="{{generateBannerUrl .BannerPath}}" alt="Banner for {{.Title}}">
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
                            <li>Credit: {{$chart.Credit}}</li>
                            <li>Steps Type: {{$chart.StepsType}}</li>
                            <li>Difficulty: {{$chart.Difficulty}}</li>
                            <li> Meter: {{$chart.Meter}} </li>
                            <li> Radar Values </li>
                            <ul>
                                <li>Stream: {{$chart.Stream}}</li>
                                <li>Voltage: {{$chart.Voltage}}</li>
                                <li>Chaos: {{$chart.Chaos}}</li>
                                <li>Freeze: {{$chart.Freeze}}</li>
                                <li>Air: {{$chart.Air}}</li>
                            </ul>
                        </ul>
                    {{end}}
            </div>

            
        </ul>
    </div>
    
{{end}}