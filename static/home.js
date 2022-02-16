window.addEventListener("load", function () {
  console.log(document.getElementById("editor"));
  var cm = CodeMirror.fromTextArea(document.getElementById("editor"));
  // var cm = CodeMirror.fromTextArea(document.body);
  console.log(cm);
  // cm.getDoc().setValue("Hello world!");
});
