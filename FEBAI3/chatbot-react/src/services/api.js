export async function processDialog(prompt) {
    const response = await fetch("http://localhost:8080/api/dialog/process", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ prompt }),
    });
    return response.json();
  }
  
  export async function manualDialog(content) {
    const response = await fetch("http://localhost:8080/api/dialog/manual", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ content }),
    });
    return response.json();
  }
  