{{define "base"}}

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
    <br>
    <nav>
        <a href="/songs"> Home</a> | 
        <a href="/about"> About</a>
    </nav>


    
</header>

<main>
    {{template "content" .}}
</main>

<footer>
    <h1 class="centered">You have successfully reached the end of the page. Good Job!</h1>
</footer>


</body>
</html>

{{end}}