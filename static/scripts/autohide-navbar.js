var prevScrollpos = window.pageYOffset;
window.onscroll = function() {
  var currentScrollPos = window.pageYOffset;
  if (prevScrollpos > currentScrollPos) {
    document.getElementById("head").style.top = "10px";
  } else {
    document.getElementById("head").style.top = "-70px";
  }
  prevScrollpos = currentScrollPos;
} 
