document
  .getElementById("encryptionForm")
  .addEventListener("submit", async function (event) {
    event.preventDefault();

    let formData = new FormData(this);
    let operation = formData.get("operation");
    let url = operation === "encrypt" ? "/encrypt" : "/decrypt";
    let inputData = formData.get("inputData");

    try {
      let response = await fetch(url, {
        method: "POST",
        body: inputData,
        headers: {
          "Content-Type": "text/plain",
        },
      });

      if (response.ok) {
        let result = await response.text();
        document.getElementById("output").innerText = result;
      } else {
        document.getElementById("output").innerText =
          "Error: " + response.status;
      }
    } catch (error) {
      console.error("Error:", error);
      document.getElementById("output").innerText = "Error: " + error.message;
    }
  });
