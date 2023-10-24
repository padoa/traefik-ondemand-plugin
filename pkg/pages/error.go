package pages

import (
	"bytes"
	"html/template"
	"path"
)

var errorPage = `<!doctype html>
<html lang="en-US">

<head>
  <title>Ondemand - Error</title>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />

  <meta http-equiv="refresh" content="5" />

  <link rel="shortcut icon" href="https://docs.traefik.io/assets/images/logo-traefik-proxy-logo.svg" />
  <link rel="preconnect" href="https://fonts.gstatic.com/">

  <style type="text/css">
    :root {
      --color-Rose: #ff99a5;
      --color-Jaune: #ffcc01;
      --color-Vert: #00cc99;
      --color-Raven: #0046da;
      --color-Raven-shadow: #0046da66;
      --color-Beige: #f6ecdf;
      --color-Wayne6: #001440;
      --translate-size: 7px;
    }
    html {
      background-color: var(--color-Wayne6);
    }
    body {
      height: 100vh;
      width: 100vw;
      font-family: 'Work Sans', sans-serif;
      margin: 0;
      padding: 0;
    }
    .u-flex-center {
      display: flex;
      justify-content: center;
      align-items: center;
      overflow-y: auto;
    }
    .cluster {
      width: fit-content;
      margin: auto;
      margin-top: 30px;
      margin-bottom: 20px;
      border-radius: 24px;
      padding: 24px 46px 1px;
      background-color: var(--color-Beige);
      position: relative;
      transition: 300ms ease-in-out;
      min-width: 200px;
      max-width: 70%;
    }
    .cluster:before {
      content: '';
      background-color: var(--color-Vert);
      border-top-left-radius: 24px;
      border-bottom-left-radius: 24px;
      position: absolute;
      bottom: 0;
      left: 0;
      top: 0;
      width: 20px;
    }
    .cluster:after {
      content: '';
      background-color: var(--color-Rose);
      border-top-right-radius: 24px;
      border-bottom-right-radius: 24px;
      position: absolute;
      bottom: 0;
      right: 0;
      top: 0;
      width: 20px;
    }
    .cluster:hover {
      transform: translateY(calc(-1 * var(--translate-size)));
      box-shadow: 0 15px 0 0 var(--color-Raven);
    }
    .title {
      margin-top: 8px;
      margin-bottom: 24px;
      font-weight: 600;
      font-size: 22px;
    }
    .title.small {
      font-size: 14px;
    }
    .subtitle {
      font-weight: 600;
      font-size: 18px;
      position: relative;
    }
    .subtitle:after {
      background-color: var(--color-Jaune);
      height: 5px;
      bottom: -3px;
      content: '';
      left: 0;
      position: absolute;
      right: 0;
      transform: scaleX(0);
      transform-origin: 100% 50%;
      transition: transform 300ms ease-in-out;
    }
    .cluster:hover .subtitle:after {
      transform: scaleX(1);
      transform-origin: 0 50%;
    }
    .code {
      font-family: 'Courier New', Courier, monospace;
      background-color: var(--color-Wayne6);
      color: var(--color-Rose);
      padding: 24px;
      font-size: 16px;
      transition: 300ms ease-in-out;
    }
    .cluster:hover .code {
      transform: translate(calc(.5 * var(--translate-size)), calc(-.5 * var(--translate-size)));
      box-shadow: -10px 10px 0 0 var(--color-Jaune);
    }
    
    .footer {
      position: absolute;
      bottom: 0px;
    }
    .footer>a {
      text-decoration: none;
      color: var(--color-Beige);
      transition: 500ms;
      opacity: .4;
    }
    .footer>a:hover {
      opacity: 1;
    }

    .copyright {
      opacity: .3;
      position: fixed;
      bottom: 15px;
      right: 120px;
      color: var(--color-Beige);
      transition-duration: 250ms;
      transform: scale(0.7);
    }
    .copyright:hover {
      opacity: 1;
      transform: scale(1);
    }
    .copyright:before,
    .copyright:after {
      opacity: 0;
      position: absolute;
      transition-duration: 250ms;
      white-space: nowrap;
    }
    .copyright:before {
      content: "Designed with";
      left: -40px;
    }
    .copyright:after {
      content: "by Staylix";
      right: -35px;
    }
    .copyright:hover:before,
    .copyright:hover:after {
      opacity: 1;
      transform: translateX(0);
    }
    .copyright:hover:before {
      left: -110px;
    }
    .copyright:hover:after {
      right: -78px;
    }
    .heart {
      background-color: var(--color-Rose);
      display: inline-block;
      height: 14px;
      position: relative;
      top: 0;
      transform: rotate(-45deg);
      width: 15px;
      border-radius: 2px;
    }
    .heart:before,
    .heart:after {
      content: "";
      background-color: var(--color-Rose);
      border-radius: 50%;
      height: 14px;
      position: absolute;
      width: 14px;
    }
    .heart:before {
      top: -6px;
      left: 0;
    }
    .heart:after {
      left: 6px;
      top: 0;
    }
  </style>
</head>


<body class="u-flex-center">
  <div class="cluster">
    <div>
      <span class="subtitle">Nom de votre stack</span>
      <div class="title">{{ .Name }}</div>
    </div>
    <div>
      <span class="subtitle">Erreur :(</span>
      <div class="title small">
        Une erreur a eu lieu pendant le réveil de votre stack.<br/>
        Contactez l'équipe SRE (#team_sre).
      </div>
    </div>
    <div class="title code">
      {{ .Error }}
    </div>
  </div>

  <div class="copyright">
    <div class="heart"></div>
  </div>

  <footer class="footer title small">
    <a href="https://github.com/acouvreur/traefik-ondemand-plugin"
      target="_blank">acouvreur/traefik-ondemand-plugin</a>
  </footer>
</body>

</html>`

type ErrorData struct {
	Name  string
	Error string
}

func GetErrorPage(template_path string, name string, e string) string {
	var tpl *template.Template
	var err error
	if template_path != "" {
		tpl, err = template.New(path.Base(template_path)).ParseFiles(template_path)
	} else {
		tpl, err = template.New("error").Parse(errorPage)
	}
	if err != nil {
		return err.Error()
	}

	b := bytes.Buffer{}
	err = tpl.Execute(&b, ErrorData{
		Name:  name,
		Error: e,
	})
	if err != nil {
		return err.Error()
	}

	return b.String()
}
