<!DOCTYPE html>
<html>
  <head>
      <meta charset="utf-8">
      
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css">
      <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>     
<style>
.container {
    padding: 16px;
}
</style>
<form>
  <div class="container">
    <h1>Docker Names List</h1>
        {{ range $key, $value := .Names }}
        <a href="/docker/config/{{index $.ID $key}}">NAME: {{ $value}}</a>
    <br>
        {{end}}
    </div>
</form>
