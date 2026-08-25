const socket = new WebSocket("ws://localhost:8082/ws")

const card = document.getElementById("card")
const status = document.getElementById("status")
const el = (id) => document.getElementById(id)

const STAT_KEYS = ["strength", "speed", "technique", "charisma", "stamina"]

socket.onopen = () => (status.textContent = "Connected — waiting for the deal")
socket.onclose = () => (status.textContent = "Disconnected")
socket.onerror = () => (status.textContent = "Connection failed")

socket.onmessage = (e) => {
  const msg = JSON.parse(e.data)
  if (msg.type !== "hand" || !msg.cards?.length) return
  renderCard(msg.cards[0])
  status.textContent = `Card 1 of ${msg.cards.length} in hand`
}

function initials(name) {
  const words = name.trim().split(/[\s-]+/)
  if (words.length > 1) return (words[0][0] + words[1][0]).toUpperCase()
  return name.slice(0, 2).toUpperCase()
}

function imperialHeight(cm) {
  const inches = Math.round(cm / 2.54)
  return `${Math.floor(inches / 12)}'${inches % 12}"`
}

function renderCard(c) {
  card.dataset.brand = c.brand.toLowerCase()
  card.dataset.alignment = c.alignment.toLowerCase()

  el("ovr").textContent = c.card_stats.overall
  el("brand").textContent = c.brand
  el("alignment").textContent = c.alignment
  el("monogram").textContent = initials(c.name)
  el("name").textContent = c.name
  el("realName").textContent = c.real_name
  el("hometown").textContent = c.hometown

  el("height").textContent = `${c.height_cm} CM`
  el("heightImp").textContent = imperialHeight(c.height_cm)
  el("weight").textContent = `${c.weight_kg} KG`
  el("weightImp").textContent = `${Math.round(c.weight_kg * 2.205)} LB`
  el("debut").textContent = c.debut_year

  el("finisher").textContent = c.finisher
  el("signature").textContent = c.signature_move

  const rows = document.querySelectorAll(".stat")
  rows.forEach((row) => {
    const value = c.card_stats[row.dataset.stat]
    row.querySelector(".stat__value").textContent = value
    row.style.setProperty("--v", 0)
  })

  requestAnimationFrame(() => {
    rows.forEach((row) => row.style.setProperty("--v", c.card_stats[row.dataset.stat]))
  })

  const belt = el("belt")
  const held = !c.current_title.startsWith("None")
  belt.classList.toggle("is-none", !held)
  el("title").textContent = held ? c.current_title : "No active title"


  card.classList.remove("deal-in")
  void card.offsetWidth
  card.classList.add("deal-in")
}

card.addEventListener("pointermove", (e) => {
  const box = card.getBoundingClientRect()
  const x = (e.clientX - box.left) / box.width
  const y = (e.clientY - box.top) / box.height
  card.style.setProperty("--mx", `${x * 100}%`)
  card.style.setProperty("--my", `${y * 100}%`)
  card.style.setProperty("--rx", `${(x - 0.5) * 16}deg`)
  card.style.setProperty("--ry", `${(0.5 - y) * 16}deg`)
})

card.addEventListener("pointerleave", () => {
  card.style.setProperty("--mx", "50%")
  card.style.setProperty("--my", "50%")
  card.style.setProperty("--rx", "0deg")
  card.style.setProperty("--ry", "0deg")
})
