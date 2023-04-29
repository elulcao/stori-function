package email

const EmailTemplate = `
{{define "email"}}
	<html>
	  <head>
	    <title>{{.Transaction.Owner}}</title>
	    <style>
	      table {
	        width: 100%;
	      }
	      td {
	        padding: 5px;
	      }
	    </style>
	  </head>
	  <body>
	    <table>
	      <tr>
	        <td style="width: 50%;">Total balance is {{printf "%.2f" .Transaction.TotalBalance}}</td>
	        <td style="width: 50%;">
						Average debit amount: {{printf "%.2f" .Transaction.DebitAmount}}
            <br>
						Average credit amount: {{printf "%.2f" .Transaction.CreditAmount}}</td>
	      </tr>
	      {{range $key, $value := .Transaction.Transactions}}
	      <tr>
	        <td>Number of transactions in {{$key}}: {{$value}}</td>
	        <td></td>
	      </tr>
	      {{end}}
	    </table>
			<br>
			<br>
	    <p>Thank you for your business!</p>
			<br>
			<br>
	    <p>Sincerely,</p>
	    <p>{{.Sender}}</p>
	    <div class="small-image-container">
	      <img src="{{(printf "%s/logo.svg" "assets") | resources.GetByPrefix}}" alt="Logo" style="max-width: 100%;">
	    </div>
	  </body>
	</html>
{{end}}
`
