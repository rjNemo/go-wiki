let quill = new Quill("#editor", {
  theme: "snow"
});

function save() {
  let text = quill.getContents();
  console.log(text);

  fetch("http://localhost:8080/editor", {
    method: "post",
    headers: { Accept: "application/json", "Content-Type": "application/json" }
  });

  console.log("send to backend");
}
