<!DOCTYPE html>

<html>
<head>
   <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="../../favicon.ico">

    <title>Home</title>

    <!-- Bootstrap core CSS -->
    <link href="http://cdn.bootcss.com/bootstrap/3.3.4/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <!--<link href="dashboard.css" rel="stylesheet">-->
</head>

<body>
<h2 class="sub-header">Home</h2>
  <div class="table-responsive">
            <table class="table table-striped">
              <thead>
                <tr>
                  <th style="width: 5%">id</th>
                  <th style="width: 15%">teama</th>
                  <th style="width: 15%">teamb</th>
                  <th>others</th>
				 <th>others</th>
                </tr>
              </thead>
              <tbody>
			{{range $k,$v:=.datas}}
                <tr>
                <td>{{$k}}</td>
           		       <td>{{$v.TeamA}}</td>
           		       <td>{{$v.TeamB}}</td>
           		       <td>{{$v.Timestamp}}</td>
						<td>{{$v.Begintime}}</td>
						<td>
                </tr>
				{{end}}
              </tbody>
			
            </table>
          </div>
<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="http://cdn.bootcss.com/jquery/1.11.2/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="http://cdn.bootcss.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
</body>

</html>
