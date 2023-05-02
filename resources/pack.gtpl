{{define "content"}}
<div>
    <h1>{{.PackName}}</h1>
    {{range .Songs}}
        <a href="/songs/{{.SongId}}">
            <div class="list-item">
                <h3 style="font-weight:bold;">{{.Title}} </h3>
                <ul>
                    <li>Pack: {{.PackName}}</li>
                    <li>Artist: {{.Artist}}</li>
                    <li>Bpms: {{range $i, $a := .Bpms}} {{$a.Value}}, {{end}}</li>
                    <li>Difficulty Meters: {{range $i, $chart := .Charts}} {{$chart.Meter}}, {{end}} </li>
                </ul>
            </div>
        </a>
        {{end}}
</div>
{{end}}