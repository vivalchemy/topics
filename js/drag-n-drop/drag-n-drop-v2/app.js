const gridItems = document.querySelectorAll(".grid-item");

gridItems.forEach((item) => {
  item.addEventListener("dragstart", dragStart);
  item.addEventListener("dragover", dragOver);
  item.addEventListener("dragleave", dragLeave);
  item.addEventListener("drop", drop);
  item.addEventListener("dragend", dragEnd);
});

let draggedItem = null;

function dragStart(e) {
  draggedItem = this;
  e.dataTransfer.effectAllowed = "move";
}

function dragOver(e) {
  e.preventDefault(); // Necessary to allow dropping
  e.dataTransfer.dropEffect = "move";
  this.classList.add("drag-over"); // Add a class to indicate a valid drop target
}

function dragLeave(e) {
  e.preventDefault();
  this.classList.remove("drag-over");
}

function drop(e) {
  e.preventDefault();
  this.classList.remove("drag-over"); // Remove the drag-over class

  // Swap the text of the dragged item and the dropped item
  [draggedItem.innerHTML, this.innerHTML] = [
    this.innerHTML,
    draggedItem.innerHTML,
  ];
}

function dragEnd() {
  gridItems.forEach((item) => {
    item.classList.remove("drag-over"); // Clean up the class after dragging ends
  });
}
