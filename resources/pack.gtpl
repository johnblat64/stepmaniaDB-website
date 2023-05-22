{{define "content"}}
<div>
    <h1>{{.PackName}}</h1>
    {{if ne .BannerPath ""}}
        <img src="{{generateBannerUrl .BannerPath}}" alt="Banner for {{.PackName}}">
    {{end}}
    <a href="{{.DownloadLink}}"> Click here to Download </a>
    {{range $i, $song := .Songs}}
        <a href="/songs/{{.SongId}}">
            <div class="list-item">
                <img src="{{generateBannerUrl .BannerPath}}" alt="Banner for {{.Title}}">
                <h3 style="font-weight:bold;">{{.Title}} </h3>
                <ul>
                    <li>Pack: {{.PackName}}</li>
                    <li>Artist: {{$song.Artist}}</li>
                    <li>Bpms: {{range $i, $bpm := $song.Bpms}} {{$bpm.Value}}, {{end}}</li>
                    <li>Difficulty Meters: 
                        {{range $i, $chart := $song.Charts}} 
                            {{$chart.Meter}}, 
                        {{end}} 
                    </li>
                </ul>
            </div>
        </a>
        {{end}}
</div>
{{end}}