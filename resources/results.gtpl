<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/resources/p5.css">
    <title>StepmaniaDB</title>
</head>
<body>

<header>
    <span class="title">    
        <h1 style="display:inline;">www.StepmaniaDB<a href="https://youtu.be/ruSjI7r1sO0?t=294">.</a>com</h1>
        <img style="width:80px;" src="/resources/ddr.gif" alt="ddr">
    </span>

    <div class="intro">
        <small></small>
        <br>
        <p> This is a fun little site I made that can search through a database of stepmania song and chart metadata. Try out all the different form fields until you find a combination that is right for you! I personally enjoy the Steps Type search field as it allows me to search for specific games and pad configurations. I hope you get use out of all the fields to find songs you like </p>
        <p> If there is a pack that is not on here, please send me the download link and I'll try and get it added haha. Message me on this discord server: <a style="color:blue" href="https://discord.gg/Pj2PCJnggT"> https://discord.gg/Pj2PCJnggT </a></p> <br>
    </div>
</header>

<main>

    <div class="searchzone-header">
        <h2>Search Zone</h2>
    </div>
    <div class="flex-container">


        <div class="left-search-form" >
            <h2>Search Fields</h3>
            <form  action="/songs" method="GET">
                <label for="title">Song Title</label> 
                <input type="text" name="title" id="title"  value="{{.SongResultsModel.SearchParameters.Title}}"> 

                <label for="artist" >Artist</label> 
                <input type="text" name="artist" id="artist" value="{{.SongResultsModel.SearchParameters.Artist}}">
                <!-- text inputs for credit, pack name, time signature (in the form numerator/denominator)-->
                <label for="credit" >Credit (Chart Author)</label> 
                <input type="text" name="credit" id="credit" value="{{.SongResultsModel.SearchParameters.Credit}}">

                <label for="pack" >Pack Name</label>
                <input type="text" name="pack" id="pack" value="{{.SongResultsModel.SearchParameters.Pack}}">
                <!-- The HTML input types are:-->
                <!-- text, number, date, time, datetime-local, month, week, email, url, search, tel, and password.-->

                <label for="timeSignatureNumerator">Time Signature</label>
                <input type="number" name="timeSignatureNumerator" id="timeSignatureNumerator" value="4">
                <input type="number" name="timeSignatureDenominator" id="timeSignatureDenominator" value="4">

                <label for="bpmMin">BPM</label> 
                <label for="bpmMin">Min</label> <input type="number" name="bpmMin" id="bpmMin" value="0">
                <label for="bpmMax">Max</label> <input type="number" name="bpmMax" id="bpmMax" value="999">
                <label for="stepstype">Steps Type</label> 
                <select name="stepstype" id="stepstype">
                    <option value="">All</option>
                
                    {{range $i, $a := .SongResultsModel.StepsTypeOptions}}
                    <option value="{{$a}}" {{if eq $a $.SongResultsModel.SearchParameters.StepsType}} selected {{end}}>{{$a}}</option>
                    {{end}}

                </select><br>
                <input type="submit" value="Search">
            </form>
        </div>


        <div class="right-search-results">
            <h2>Search Results</h2>

            <div id="page-selection-top">
                <p>Number of Results: {{.SongResultsModel.TotalSongsCount}}</p>
                {{if ge .SongResultsModel.PreviousPage 1}}
                    <a href="songs{{.SongResultsModel.SearchParameters.AsQueryString}}&page={{.SongResultsModel.PreviousPage}}">&lt;===</a>
                {{end}}
                
                {{.SongResultsModel.Page}}  

                {{if .SongResultsModel.HasNextPage}}
                    <a href="songs{{.SongResultsModel.SearchParameters.AsQueryString}}&page={{.SongResultsModel.NextPage}}">===&gt;</a>
                {{end}}
            </div>

            {{range .SongResultsModel.Songs}}
            <a href="songs/{{.SongId}}">
                <article class="list-item">
                    <h3>{{.Title}} </h3>
                    <ul>
                        <li>Artist: {{.Artist}}</li>
                        <li> Bpms: {{range $i, $a := .Bpms}} {{$a.Value}}, {{end}}</li>
                    </ul>
                </article>
            </a>
            {{end}}

            <div id="page-selection-bottom">
                <p>Number of Results: {{.SongResultsModel.TotalSongsCount}}</p>
                {{if ge .SongResultsModel.PreviousPage 1}}
                    <a href="songs{{.SongResultsModel.SearchParameters.AsQueryString}}&page={{.SongResultsModel.PreviousPage}}">&lt;===</a>
                {{end}}
                
                {{.SongResultsModel.Page}}  
                
                {{if .SongResultsModel.HasNextPage}}
                    <a href="songs{{.SongResultsModel.SearchParameters.AsQueryString}}&page={{.SongResultsModel.NextPage}}">===&gt;</a>
                {{end}}
            </div>
        </div>
        


    </div>

</main>


<footer>
    <h1 class="centered">You have successfully reached the end of the page</h1>
    <!-- Add any additional footer content here -->
</footer>



</body>
</html>
