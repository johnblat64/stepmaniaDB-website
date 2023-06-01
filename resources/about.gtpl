{{define "content"}}

<h2>What the heck is this site?</h2>

    <p> This is a fun little site I made that can search through a database of stepmania song and chart metadata. Try out all the different form fields until you find a combination that is right for you! I personally enjoy the Steps Type search field as it allows me to search for specific games and pad configurations. I hope you get use out of all the fields to find songs you like </p>
    <p> If there is a pack that is not on here, please send me the download link and I'll try and get it added haha. Message me on this discord server: <a href="https://discord.gg/Pj2PCJnggT"> https://discord.gg/Pj2PCJnggT </a></p> <br>


<h2> Why did I make this site?</h2>
I made this site because there wasn't a site that met the following critieria:
<ul>
<li>Search by fields such as BPMs, Games, Pad configurations</li> 
<li>Allowed for ways for users to contribute to the database</li> 
<li>Maintain a centralized repository of songs with download links</li> 
</ul>

stepmaniaonline.net is a great site, but it doesn't allow for searching by fields such as BPMs, Games, Pad configurations. It also doesn't explain how users to contribute to the database. 

<br>
<h2>Technical stuff:</h2>

<!--make the lines bellow into an unordered list-->
<ul>
<li>I wrote a parser that can parse stepmania song files to extract metadata and load it into a database and stores. </li>
<li>I wrote a backend web api that can query the database and return the results as json. </li>
<li>I wrote a frontend site that calls the backend api for display to the user.</li>
</ul>




The reason i separated these is so that user developers can use the backend api to create their own frontend sites, and also potentially create an integration with a StepMania aoo that can download songs from the site directly into the game. 

{{end}}