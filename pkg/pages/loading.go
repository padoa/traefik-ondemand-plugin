package pages

import (
	"bytes"
	"path"

	"fmt"
	"html/template"
	"math"
	"time"
)

var loadingPage = `<!doctype html>
<html lang="en-US">

<head>
  <title>Ondemand - Loading</title>
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
    }
    .container {
      height: 100%;
      width: 100%;
      overflow-y: auto;
    }
    .col-container {
      display: flex;
      align-items: center;
      flex: 1;
      min-height: 0px;
    }
    .col {
      display: flex;
      flex-direction: column;
      height: 100%;
      width: 50%;
      align-items: center;
      justify-content: center;
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

    @keyframes spinner-transform {
      0% { transform: translateX(0px) rotate3d(0); animation-timing-function: ease-in-out }
      19% { transform: rotate3d(0, 1, 0, 180deg); transform-origin: right 0; animation-timing-function: steps(1, start);}
      20% { transform: translateX(100px) rotate(0deg); animation-timing-function: ease-in }
      30% { transform: translateX(100px) rotate(90deg); transform-origin: bottom left; animation-timing-function: ease-out; }
      35% { transform: translateX(100px) rotate(70deg); transform-origin: bottom left; animation-timing-function: ease-in; }
      40% { transform: translateX(100px) translateY(0px) rotate(90deg) scale(1); transform-origin: bottom left; }
      55% { transform: translateX(100px) translateY(0px) rotate(500deg) scale(0); transform-origin: bottom right; animation-timing-function: step(1); }
      56% { transform: translateX(0px) translateY(100px) rotate(0deg) scale(0, 1); transform-origin: left; }
      65% { transform: translateX(0px) translateY(100px) rotate(0deg) scale(1) ; transform-origin: left; }
      70% { transform: translateX(0px) translateY(130px) rotate(0deg) scale(1) ; transform-origin: left; animation-timing-function: cubic-bezier(.3,.2,.09,1.06); }
      75% { transform: translateY(-400px); animation-timing-function: cubic-bezier(.38,0,1,.79); }
      81% { transform: translateY(0px); animation-timing-function: cubic-bezier(0,.39,.79,.99); }
      84% { transform: translateY(-60px); animation-timing-function: cubic-bezier(.38,0,1,.79); }
      87% { transform: translateY(0px) scale(1); }
      91% { transform: scale(0.1); }
      95% { transform: scale(1.6); }
      100% { transform: scale(1); }
    }
    @keyframes spinner-radius {
      0% { border-radius: 0; }

      30% { border-radius: 48px; border-bottom-left-radius: 0px; }

      56% { border-radius: 0; }

      65% { border-radius: 0; animation-timing-function: ease; }
      73% { border-radius: 100%; }

      92% { border-radius: 100%; }
      93% { border-radius: 0; }

    }
    @keyframes spinner-colors {
      0% { background-color: var(--color-Raven); }

      5% { background-color: var(--color-Vert); }

      30% { background-color: var(--color-Jaune); }

      56% { background-color: var(--color-Rose); }

      91% { background-color: var(--color-Raven); }
    }
    .spinner {
      width: 100px;
      height: 100px;
      animation: 5s ease          infinite spinner-transform,
                5s steps(1, end) infinite spinner-radius,
                5s steps(1, end) infinite spinner-colors;
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
  <div class="col-container">
    <div class="col">
      <div class="container">
        <div class="cluster">
          <div>
            <span class="subtitle">Nom de votre stack</span>
            <div class="title">{{ .Name }}</div>
          </div>
          <div>
            <span class="subtitle">Réveil en cours...</span>
            <div class="title small">
              Votre stack est en train de se réveiller, veuillez patienter quelques minutes...<br>
              En cas d'attente de plus de 15 minutes, contactez l'équipe SRE (#team_sre).
            </div>
          </div>
          <div>
            <span class="subtitle">Arrêt automatique</span>
            <div class="title small">
              Votre stack sera automatiquement arrêtée après {{ .Timeout }} d'inactivité.
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="col">
      <div class="spinner"></div>
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

type LoadingData struct {
	Name    string
	Timeout string
}

func GetLoadingPage(template_path string, name string, timeout time.Duration) string {
	var tpl *template.Template
	var err error
	if template_path != "" {
		tpl, err = template.New(path.Base(template_path)).ParseFiles(template_path)
	} else {
		tpl, err = template.New("loading").Parse(loadingPage)
	}
	if err != nil {
		return err.Error()
	}

	b := bytes.Buffer{}
	err = tpl.Execute(&b, LoadingData{
		Name:    name,
		Timeout: humanizeDuration(timeout),
	})
	if err != nil {
		return err.Error()
	}

	return b.String()
}

// humanizeDuration humanizes time.Duration output to a meaningful value,
// golang's default ``time.Duration`` output is badly formatted and unreadable.
func humanizeDuration(duration time.Duration) string {
	if duration.Seconds() < 60.0 {
		return fmt.Sprintf("%d secondes", int64(duration.Seconds()))
	}
	if duration.Minutes() < 60.0 {
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		if remainingSeconds > 0 {
			return fmt.Sprintf("%d minutes %d secondes", int64(duration.Minutes()), int64(remainingSeconds))
		}
		return fmt.Sprintf("%d minutes", int64(duration.Minutes()))
	}
	if duration.Hours() < 24.0 {
		remainingMinutes := math.Mod(duration.Minutes(), 60)
		remainingSeconds := math.Mod(duration.Seconds(), 60)

		if remainingMinutes > 0 {
			if remainingSeconds > 0 {
				return fmt.Sprintf("%d heures %d minutes %d secondes", int64(duration.Hours()), int64(remainingMinutes), int64(remainingSeconds))
			}
			return fmt.Sprintf("%d heures %d minutes", int64(duration.Hours()), int64(remainingMinutes))
		}
		return fmt.Sprintf("%d heures", int64(duration.Hours()))
	}
	remainingHours := math.Mod(duration.Hours(), 24)
	remainingMinutes := math.Mod(duration.Minutes(), 60)
	remainingSeconds := math.Mod(duration.Seconds(), 60)
	return fmt.Sprintf("%d jours %d heures %d minutes %d secondes",
		int64(duration.Hours()/24), int64(remainingHours),
		int64(remainingMinutes), int64(remainingSeconds))
}
