//function reload() {
//  setInterval(function () {
//    location.reload();
//  }, 3000);
//}
//
//reload();

//-------------------------------------
//
let containers = document.querySelectorAll(".container");

containers.forEach((container) => {
  // container which is dragged
  container.addEventListener("dragstart", (e) => {
    e.target.classList.add("dragged-item");
    e.dataTransfer.setData("text/plain", e.target.innerText);
  });

  // container which is dragged
  container.addEventListener("dragend", (e) => {
    e.target.classList.remove("dragged-item");
  });

  // container on which the item is dragged
  container.addEventListener("dragover", (e) => {
    e.preventDefault();
    e.target.classList.add("dragged-over");
  });

  // container from which the item was dragged
  container.addEventListener("dragleave", (e) => {
    e.preventDefault();
    e.target.classList.remove("dragged-over");
  });

  container.addEventListener("drop", (e) => {
    e.preventDefault();

    let dropLocation = e.target;
    dropLocation.classList.remove("dragged-over");
    let draggedItem = e.dataTransfer.getData("text/plain");

    dropLocation.innerText = "nothing";
    setTimeout(() => {
      dropLocation.innerText = draggedItem;
      dropLocation.classList.add("container");
      e.target.appendChild(dropLocation);
    }, 1000);
  });
});
