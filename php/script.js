console.log("Hello script");

var xhr = new XMLHttpRequest();
var url = 'http://localhost:18888/';

const handler = () => {
  console.log(xhr.responseText);
}

const getRequest = () => {
  if(xhr) {
    xhr.open('GET', url);
    xhr.onreadystatechange = handler;
    xhr.send();
  }
}

const postRequest = () => {
  if(xhr) {
    xhr.open('POST', url);
    xhr.onreadystatechange = handler;
    xhr.send();
  }
}

const postJsonRequest = () => {
  if(xhr) {
    xhr.open('POST', url);
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.setRequestHeader('X-CORS-Sample', 'foo')
    xhr.onreadystatechange = handler;
    xhr.send();
  }
}

const getCookieRequest = () => {
  if(xhr) {
    xhr.open('GET', url);
    xhr.withCredentials = true;
    xhr.onreadystatechange = handler;
    xhr.send();
  }
}

const setCookie = () => {
  document.cookie = 'hoge=123'
}

document.addEventListener('DOMContentLoaded', () => {
  setCookie()

  // getRequest()
  // postRequest()
  // postJsonRequest()
  getCookieRequest()
})
