package bindata_files

import (
	"time"

	"ktkr.us/pkg/vfs/bindata"
)

func init() {
	bindata.RegisterFile("templates/config.tmpl", time.Unix(1440739326, 0), []byte("\x7b\x7b define \"config\" \x7d\x7d\n  <section id=\"section-overview\" class=\"floating-section\">\n    \x7b\x7b template \"overview\" . \x7d\x7d\n  </section>\n  <section id=\"section-config\" class=\"floating-section\">\n    <h1>Configuration</h1>\n    <form id=\"config\">\n      <div id=\"help\">?</div>\n      <div class=\"box\" id=\"host-box\" data-tooltip=\"This is the base host and path that the URL will be appended to and returned when \x79ou upload a file. Leave it blank to use whichever host the uploader is accessed from.\" data-tt-pos=\"top\">\n        <label for=\"host\">Base path</label>\n        <input t\x79pe=\"text\" id=\"host\" name=\"host\" value=\"\x7b\x7b .Conf.Host \x7d\x7d\" placeholder=\"i.example.com\">\n      </div>\n      <div class=\"box\" id=\"port-box\" data-tooltip=\"This is the port that the server will listen on. You can set it to something other than 80 and set up a prox\x79 with \x79our existing web server.\" data-tt-pos=\"right\">\n        <label for=\"port\">Port</label>\n        <input t\x79pe=\"text\" id=\"port\" name=\"port\" value=\"\x7b\x7b .Conf.Port \x7d\x7d\" placeholder=\"80\">\n      </div>\n      <div class=\"box col3\" id=\"max-age-box\" data-tooltip=\"This is the maximum age of an upload in da\x79s. Uploads will be automaticall\x79 pruned when the\x79 reach this age. If left blank or \xe2\x89\xa40, no pruning will happen.\" data-tt-pos=\"left\">\n        <label for=\"max-age\">Max age (da\x79s)</label>\n        <input t\x79pe=\"number\" id=\"max-age\" name=\"max-age\" value=\"\x7b\x7b .Conf.Age \x7d\x7d\" min=\"0\">\n      </div>\n      <div class=\"box col3\" id=\"max-si\x7ae-box\" data-tooltip=\"This is the maximum si\x7ae of the uploads folder in megab\x79tes. The oldest uploads will be pruned whenever the si\x7ae reaches this value. If left blank or \xe2\x89\xa40, no pruning will happen.\" data-tt-pos=\"right\">\n        <label for=\"max-si\x7ae\">Max si\x7ae (MB)</label>\n        <input t\x79pe=\"number\" id=\"max-si\x7ae\" name=\"max-si\x7ae\" value=\"\x7b\x7b .Conf.Si\x7ae \x7d\x7d\" min=\"0\">\n      </div>\n      <div class=\"box col3\" id=\"hash-len-box\" data-tooltip=\"The length of the ID associated with each file. In general, longer IDs means less collision. Maximum length is 64.\" data-tt-pos=\"right\">\n        <label for=\"hash-len\">ID si\x7ae</label>\n        <input t\x79pe=\"number\" id=\"hash-len\" name=\"hash-len\" value=\"\x7b\x7b .Conf.HashLen \x7d\x7d\" min=\"1\" max=\"64\">\n      </div>\n      <div class=\"box\" id=\"director\x79-box\" data-tooltip=\"This is the director\x79 on the server in which the uploaded files will reside.\" data-tt-pos=\"left\">\n        <label for=\"director\x79\">Director\x79</label>\n        <input t\x79pe=\"text\" id=\"director\x79\" name=\"director\x79\" value=\"\x7b\x7b .Conf.Director\x79 \x7d\x7d\" placeholder=\"/home/user/uploads\">\n      </div>\n      <div class=\"box checkbox\" data-tooltip=\"Enable to append the original file extension to returned links.\" data-tt-pos=\"left\">\n        <input t\x79pe=\"checkbox\" id=\"append-ext\" name=\"append-ext\"\x7b\x7b if .Conf.AppendExt \x7d\x7dchecked\x7b\x7b end \x7d\x7d>\n        <label for=\"append-ext\">Append file extensions</label>\n      </div>\n      <div class=\"box checkbox\" data-tooltip=\"Enable to allow uploads to show Twitter cards with file previews if applicable.\" data-tt-pos=\"left\">\n        <input t\x79pe=\"checkbox\" id=\"twitter-card\" name=\"twitter-card\"\x7b\x7b if .Conf.TwitterCardEnable \x7d\x7dchecked\x7b\x7b end \x7d\x7d>\n        <label for=\"twitter-card\">Enable Twitter Cards</label>\n        <div id=\"twitter-card--hidden\">\n          <label for=\"twitter-handle\">Twitter handle</label>\n          <input t\x79pe=\"text\" id=\"twitter-handle\" name=\"twitter-handle\" value=\"\x7b\x7b .Conf.TwitterHandle \x7d\x7d\" required placeholder=\"@handle\">\n        </div>\n      </div>\n      <div class=\"box\" id=\"newpass-box\" data-tooltip=\"Enter a new password here to change \x79our password.\" data-tt-pos=\"right\">\n        <label for=\"newpass\">New password</label>\n        <input t\x79pe=\"password\" id=\"newpass\" name=\"newpass\" placeholder=\"\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\">\n      </div>\n      <hr>\n      <div class=\"box\" id=\"password-box\" data-tooltip=\"Whatever changes \x79ou make, enter \x79our current password here to make sure that \x79ou're \x79ou. If \x79ou haven't set a password \x79et, though, \x79ou don't have to fill it out.\" data-tt-pos=\"left\">\n        <label for=\"password\">Current password</label>\n        <input t\x79pe=\"password\" id=\"password\" name=\"password\" required placeholder=\"Required\">\n      </div>\n      <hr>\n      <button id=\"submit\" t\x79pe=\"button\">Update configuration</button>\n    </form>\n  </section>\n  <script>\n    var oldMaxSi\x7ae, oldMaxAge;\n\n    function reloadOverview() \x7b\n      var x = new XMLHttpRequest();\n      x.open('GET', '/config/overview', true);\n      x.addEventListener('load', function(e) \x7b\n        if (e.target.status === 200) \x7b\n          $('#section-overview').innerHTML = e.target.response;\n        \x7d\n      \x7d);\n      x.send();\n    \x7d\n\n    window.addEventListener('load', function() \x7b\n      var buttons = $$('button');\n      oldMaxSi\x7ae  = parseInt($('#max-si\x7ae').value);\n      oldMaxAge   = parseInt($('#max-age').value);\n\n\n      $('#submit').addEventListener('click', function() \x7b\n        for (var i = 0, button; button = buttons[i]; i++) button.setAttribute('disabled', true);\n        var maxSi\x7ae = parseInt($('#max-si\x7ae').value);\n        var maxAge  = parseInt($('#max-age').value);\n        var delta   = 0;\n\n        var f = function(url, val) \x7b\n          var fd = new FormData();\n          fd.append('N', val);\n          var x = new XMLHttpRequest();\n          x.open('POST', url, false);\n          x.send(fd);\n\n          if (x.status == 200) \x7b\n            var n = JSON.parse(x.response).N;\n            if (n > delta) delta = n;\n            return true;\n          \x7d else \x7b\n            var err = JSON.parse(x.response);\n            showMessage($('#section-config'), 'Server error: ' + err.Err + ' (' + x.status + ')', 'bad');\n            return false;\n          \x7d\n        \x7d\n\n        if (maxSi\x7ae > 0 && (oldMaxSi\x7ae == 0 \x7c\x7c maxSi\x7ae < oldMaxSi\x7ae)) \x7b\n          if (!f('/config/si\x7ae', maxSi\x7ae)) return;\n        \x7d\n        if (maxAge > 0 && (oldMaxAge == 0 \x7c\x7c maxAge < oldMaxAge)) \x7b\n          if (!f('/config/age', maxAge)) return;\n        \x7d\n        if (delta > 0) \x7b\n          if (!confirm('Changes made to age or si\x7ae limits mean that ' + delta + ' old file(s) will be pruned. Continue?')) \x7b\n            return;\n          \x7d\n        \x7d\n\n        oldMaxAge = maxAge;\n        oldMaxSi\x7ae = maxSi\x7ae;\n\n        var host   = $('#host');\n        host.value = host.value.replace(/\\w+:\\/\\//, '');\n        var fd     = new FormData($('#config'));\n        var x      = new XMLHttpRequest();\n\n        x.addEventListener('load', function(e) \x7b\n          $('#password').value = '';\n          for (var i = 0, button; button = buttons[i]; i++) button.removeAttribute('disabled');\n          if (e.target.status === 204) \x7b\n            showMessage($('#section-config'), 'Configuration updated.', 'good');\n            $('#newpass').value = '';\n            reloadOverview();\n          \x7d else \x7b\n            var resp = JSON.parse(x.responseText);\n            if (resp != null && resp.Err != null) \x7b\n              showMessage($('#section-config'), 'Error: ' + resp.Err, 'bad');\n            \x7d else \x7b\n              showMessage($('#section-config'), 'An unknown error occurred (status ' + e.target.status + ')', 'bad');\n            \x7d\n          \x7d\n        \x7d, false);\n\n        x.open('POST', '/config', true);\n        x.send(fd);\n      \x7d, false);\n    \x7d, false);\n  </script>\n\x7b\x7b end \x7d\x7d\n\n\x7b\x7b define \"overview\" \x7d\x7d\n    <h1>Overview</h1>\n    <p><strong><a href=\"/histor\x79/0\">\x7b\x7b .NumUploads \x7d\x7d upload\x7b\x7b if ne .NumUploads 1 \x7d\x7ds\x7b\x7b end \x7d\x7d</a></strong> totalling \x7b\x7b .UploadsSi\x7ae \x7d\x7d. (<a href=\"javascript:purgeAll()\">purge</a>)</p>\n    <p>Thumbnail cache is \x7b\x7b .ThumbsSi\x7ae \x7d\x7d. (<a href=\"javascript:purgeThumbs()\">purge</a>)</p>\n\x7b\x7b end \x7d\x7d\n"))
	bindata.RegisterFile("templates/errors/errors.tmpl", time.Unix(1440218376, 0), []byte("\x7b\x7b define \"400\" \x7d\x7d<!doct\x79pe html>\n<html>\n  <head>\n    <title>400</title>\n    <link rel=\"st\x79lesheet\" href=\"/static/st\x79le.css\">\n  </head>\n  <bod\x79>\n    <div class=\"error\">\n      <h1>You're doing it wrong.</h1>\n      \x7b\x7b if .Err \x7d\x7d<p>\x7b\x7b .Err \x7d\x7d</p>\x7b\x7b end \x7d\x7d\n    </div>\n  </bod\x79>\n</html>\n\x7b\x7b end \x7d\x7d\n\x7b\x7b define \"404\" \x7d\x7d<!doct\x79pe html>\n<html>\n  <head>\n    <title>404</title>\n    <link rel=\"st\x79lesheet\" href=\"/static/st\x79le.css\">\n  </head>\n  <bod\x79>\n    <div class=\"error\">\n      <h1>This isn't the page \x79ou're looking for.</h1>\n    </div>\n  </bod\x79>\n</html>\n\x7b\x7b end \x7d\x7d\n\x7b\x7b define \"500\" \x7d\x7d<!doct\x79pe html>\n<html>\n  <head>\n    <title>500</title>\n    <link rel=\"st\x79lesheet\" href=\"/static/st\x79le.css\">\n  </head>\n  <bod\x79>\n    <div class=\"error\">\n      <h1>Something went wrong.</h1>\n      \x7b\x7b if .Err \x7d\x7d<p>\x7b\x7b .Err \x7d\x7d</p>\x7b\x7b end \x7d\x7d\n    </div>\n  </bod\x79>\n</html>\n\x7b\x7b end \x7d\x7d\n"))
	bindata.RegisterFile("templates/history.tmpl", time.Unix(1440740011, 0), []byte("\x7b\x7b define \"histor\x79\" \x7d\x7d\n<section id=\"histor\x79\">\n  \x7b\x7b if len .List \x7c lt 25 \x7d\x7d\x7b\x7b template \"pagination\" . \x7d\x7d\x7b\x7b end \x7d\x7d\n  <ul>\n    \x7b\x7b range .List \x7d\x7d\n    <li class=\"histor\x79-item\" data-id=\"\x7b\x7b .ID \x7d\x7d\">\n      <a href=\"/\x7b\x7b .ID \x7d\x7d\x7b\x7b if $.AppendExt \x7d\x7d\x7b\x7b .Ext \x7d\x7d\x7b\x7b end \x7d\x7d\" class=\"upload-link\">\x7b\x7b if .HasThumb \x7d\x7d<img src=\"/thumb/\x7b\x7b .ID \x7d\x7d.jpg\">\x7b\x7b else \x7d\x7d<img src=\"/static/file.svg\"><div class=\"file-ext-overla\x79\">\x7b\x7b .Ext \x7d\x7d</div>\x7b\x7b end \x7d\x7d</a>\n      <div class=\"histor\x79-item-name\" title=\"\x7b\x7b .Name \x7d\x7d\">\x7b\x7b .Name \x7d\x7d</div>\n      <div class=\"histor\x79-item-data\">\x7b\x7b .Si\x7ae \x7d\x7d / <span title=\"\x7b\x7b .Uploaded.Format \"2006-01-02 15:04:05 MST\" \x7d\x7d\">\x7b\x7b .Ago \x7d\x7d</span> ago</div>\n      <div class=\"histor\x79-item-data\"><a href=\"javascript:\" class=\"delete-upload\">Delete</a></div>\n    </li>\n    \x7b\x7b end \x7d\x7d\n  </ul>\n  \x7b\x7b template \"pagination\" . \x7d\x7d\n</section>\n<script>\n  window.addEventListener('DOMContentLoaded', function() \x7b\n    var items = $$('.histor\x79-item');\n    for (var i = 0, item; item = items[i]; i++) \x7b\n      (function(item) \x7b\n        item.quer\x79Selector('a.delete-upload').addEventListener('click', function() \x7b\n          item.st\x79le.opacit\x79 = '0.5';\n          var x = new XMLHttpRequest();\n\n          x.addEventListener('load', function(e) \x7b\n            if (e.target.status == 204) \x7b\n              item.st\x79le.opacit\x79 = '0.0';\n              item.addEventListener('transitionend', function(e) \x7b\n                e.target.parentNode.removeChild(e.target);\n                window.location.reload(true);\n              \x7d, false);\n            \x7d else \x7b\n              item.st\x79le.opacit\x79 = '';\n              var resp = JSON.parse(x.responseText);\n              if (resp != null && resp.Err != null) \x7b\n                alert('error: ' + resp.Err);\n              \x7d else \x7b\n                alert('wat');\n              \x7d\n            \x7d\n          \x7d);\n\n          x.open('POST', '/delete/' + item.dataset.id, true);\n          x.send();\n        \x7d, false);\n      \x7d)(item);\n    \x7d\n  \x7d, true);\n</script>\n\x7b\x7b end \x7d\x7d\n\n\x7b\x7b define \"pagination\" \x7d\x7d\n<nav class=\"pagination\">\n  <span class=\"prevnext\x7b\x7b if gt .CurrentPage 1 \x7d\x7d active\x7b\x7b end \x7d\x7d\"><a href=\"/histor\x79/\x7b\x7b .PrevPage \x7d\x7d\">Back</a> \xe2\x80\x94</span>\n  Page \x7b\x7b .CurrentPage \x7d\x7d of \x7b\x7b .TotalPages \x7d\x7d\n  <span class=\"prevnext\x7b\x7b if ne .NextPage 0 \x7d\x7d active\x7b\x7b end \x7d\x7d\">\xe2\x80\x94 <a href=\"/histor\x79/\x7b\x7b .NextPage \x7d\x7d\">Next</a></span>\n</nav>\n\x7b\x7b end \x7d\x7d\n"))
	bindata.RegisterFile("templates/index.tmpl", time.Unix(1440218376, 0), []byte("\x7b\x7b define \"index\" \x7d\x7d\n  <section id=\"upload\" class=\"floating-section\">\n    <input t\x79pe=\"file\" id=\"picker\" name=\"picker[]\" multiple>\n    <div id=\"drop-\x7aone\">\n      <div class=\"progress-bar\"></div>\n      <div id=\"drop-\x7aone-text\">Click/tap/drop</div>\n    </div>\n    <div id=\"uploaded-urls\">\n      <ul></ul>\n    </div>\n  </section>\n  <script>window.addEventListener('load', setupUploader, false);</script>\n\x7b\x7b end \x7d\x7d\n"))
	bindata.RegisterFile("templates/layout.tmpl", time.Unix(1440218376, 0), []byte("\x7b\x7b define \"head\" \x7d\x7d\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n<link rel=\"shortcut icon\" href=\"/static/favicon.png\">\n<link rel=\"apple-touch-icon\" si\x7aes=\"76x76\" href=\"/static/airlift_76x76.png\">\n<link rel=\"apple-touch-icon\" si\x7aes=\"120x120\" href=\"/static/airlift_120x120.png\">\n<link rel=\"apple-touch-icon\" si\x7aes=\"152x152\" href=\"/static/airlift_152x152.png\">\n<link rel=\"apple-touch-icon\" si\x7aes=\"180x180\" href=\"/static/airlift_180x180.png\">\n<link rel=\"st\x79lesheet\" href=\"/static/st\x79le.css\">\n<script src=\"/static/script.js\"></script>\n\x7b\x7b end \x7d\x7d\n\x7b\x7b define \"common\" \x7d\x7d<!doct\x79pe html>\n<html>\n  <head>\n    <title>Airlift</title>\n    \x7b\x7b template \"head\" \x7d\x7d\n  </head>\n  <bod\x79>\n    <nav id=\"nav\">\n      <a href=\"/\">Upload</a> /\n      <a href=\"/histor\x79/0\">Histor\x79</a> /\n      <a href=\"/config\">Configure</a> /\n      <a href=\"/logout\">Log out</a>\n    </nav>\n    \x7b\x7b content \x7d\x7d\n  </bod\x79>\n</html>\n\x7b\x7b end \x7d\x7d\n"))
	bindata.RegisterFile("templates/login.tmpl", time.Unix(1440218376, 0), []byte("\x7b\x7b define \"login\" \x7d\x7d<!doct\x79pe html>\n<html>\n  <head>\n    <title>Log in</title>\n    \x7b\x7b template \"head\" \x7d\x7d\n  </head>\n  <bod\x79>\n    <section id=\"section-login\" class=\"floating-section\">\n      <form method=\"post\" action=\"/login\" id=\"login\">\n        \x7b\x7b if . \x7d\x7d<p id=\"message-box\" class=\"bad\">Incorrect password.</p>\x7b\x7b end \x7d\x7d\n        <label for=\"password\">Password: </label><input name=\"pass\" id=\"password\" t\x79pe=\"password\" placeholder=\"password\" autofocus required>\n        <hr>\n        <button t\x79pe=\"submit\" id=\"submit\">Log in</button>\n      </form>\n    </section>\n  </bod\x79>\n</html>\n\x7b\x7b end \x7d\x7d\n"))
	bindata.RegisterFile("templates/twitterbot.tmpl", time.Unix(1440719768, 0), []byte("\x7b\x7b define \"twitterbot\" \x7d\x7d\n<!doct\x79pe html>\n<html>\n  <head>\n    <title>\x7b\x7b .ID \x7d\x7d</title>\n    <meta name=\"twitter:card\" content=\"summar\x79_large_image\">\n    <meta name=\"twitter:site\" content=\"\x7b\x7b .Handle \x7d\x7d\">\n    <meta name=\"twitter:title\" content=\"\x7b\x7b .ID \x7d\x7d: \x7b\x7b .Name \x7d\x7d\">\n    <meta name=\"twitter:description\" content=\"\x7b\x7b .Si\x7ae \x7d\x7d / uploaded \x7b\x7b .Uploaded.Format \"2 Jan 2006 15:04\" \x7d\x7d\">\n    <meta name=\"twitter:image\" content=\"http://\x7b\x7b .Host \x7d\x7d/twitterthumb/\x7b\x7b .ID \x7d\x7d.jpg\">\n  </head>\n  <bod\x79>\n    hi twitterbot\n  </bod\x79>\n</html>\n\x7b\x7b end \x7d\x7d\n"))
}
