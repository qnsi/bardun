function mainFunc(currentDoc) {
  defineNoteLinkMode();

  cm = createCodeMirrorEditor();
  cm.getDoc().setValue(currentDoc);

  cm.getWrapperElement().addEventListener("click", handleNoteLinkClick);
}

function handleNoteLinkClick(event) {
  var clickedDiv = event.target;
  if (clickedDiv.classList[0] === "cm-noteLink") {
    var linkNameDirty = clickedDiv.textContent;
    var linkName = linkNameDirty.substr(2).slice(0, -2);
    window.location.href = "?note=" + linkName;
  }
}

function createCodeMirrorEditor() {
  var cm = CodeMirror.fromTextArea(document.getElementById("editor"), {
    lineNumbers: true,
    lineWrapping: true,
    theme: "default",
    keyMap: "vim",
    vieportMargin: 100000000,
    mode: "noteLink",
  });
  return cm;
}

function defineNoteLinkMode() {
  CodeMirror.defineMode("noteLink", function (config, parserConfig) {
    var noteLinkOverlay = {
      token: function (stream, state) {
        var ch;
        if (stream.match("[[")) {
          while ((ch = stream.next()) != null)
            if (ch == "]" && stream.next() == "]") {
              stream.eat("]");
              // Return style for lexed element
              return "noteLink";
            }
        }
        if (stream.match("[A")) {
          while ((ch = stream.next()) != null)
            if (ch == "A" && stream.next() == "]") {
              stream.eat("]");
              // Return style for lexed element
              return "ankiLink";
            }
        }
        while (stream.next() != null && !stream.match("[[", false)) {}
        return null;
      },
    };
    return CodeMirror.overlayMode(
      CodeMirror.getMode(config, parserConfig.backdrop || "markdown"),
      noteLinkOverlay
    );
  });
}
