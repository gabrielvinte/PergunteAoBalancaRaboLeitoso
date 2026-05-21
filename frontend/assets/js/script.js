const birdAudio = new Audio("assets/audio/balancarabo.mp3");
const btnConsultar = document.getElementById("btnConsultar");
const areaResultado = document.getElementById("resultadoArea");
const textoResposta = document.getElementById("textoResposta");

btnConsultar.addEventListener("click", iniciarConsulta);

async function iniciarConsulta() {
    const op1 = document.getElementById("opcao1").value.trim();
    const op2 = document.getElementById("opcao2").value.trim();

    if (!op1 || !op2) {
        return;
    }

    try {
        areaResultado.classList.add("show");
        textoResposta.innerHTML = "";

        birdAudio.currentTime = 0;
        await birdAudio.play();

        const response = await fetch("http://localhost:8080/api/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                option_a: op1, 
                option_b: op2  
            })
        });

        const data = await response.json();
        let respostaFinal = response.ok ? data.data : "Erro: " + data.message;

        const arraySpans = [];

        for (let char of respostaFinal) {
            const span = document.createElement("span");
            span.classList.add("letra");
            
            if (char === " ") {
                span.innerHTML = "&nbsp;";
            } else {
                span.textContent = char;
                span.classList.add("oculta");
                arraySpans.push(span);
            }
            
            textoResposta.appendChild(span);
        }

        if (!birdAudio.ended) {
            await new Promise(resolve => {
                birdAudio.addEventListener("ended", resolve, { once: true });
            });
        }

        arraySpans.forEach((span, index) => {
            setTimeout(() => {
                span.classList.remove("oculta");
                span.classList.add("revelada");
            }, index * 150);
        });

    } catch (err) {
        console.error(err);
        textoResposta.textContent = "Erro na comunicação.";
        areaResultado.classList.add("show");
    }
}