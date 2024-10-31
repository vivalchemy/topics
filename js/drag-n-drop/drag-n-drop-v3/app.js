const draggables = document.querySelectorAll(".draggable");
const containers = document.querySelectorAll(".container");

// draggables
draggables.forEach((draggable) => {
  draggable.addEventListener("dragstart", (e) => {
    draggable.classList.add("dragging");
    //console.log("dragstart");
  });

  draggable.addEventListener("dragend", (e) => {
    draggable.classList.remove("dragging");
    //console.log("dragend");
  });
});

// drop locations
containers.forEach((container) => {
  container.addEventListener("dragover", (e) => {
    e.preventDefault();

    // if you want to see where the drop location is
    const afterElement = getDraggableElement(container, e.clientY);
    //console.log(afterElement);
    const draggable = document.querySelector(".dragging");
    if (!afterElement) {
      container.appendChild(draggable);
    } else {
      container.insertBefore(draggable, afterElement);
    }
    //console.log("dragover");
  });

  container.addEventListener("drop", (e) => {
    e.preventDefault();
    // if you don't want to see where the drop location is
    //const afterElement = getDraggableElement(container, e.clientY);
    ////console.log(afterElement);
    //const draggable = document.querySelector(".dragging");
    //if (!afterElement) {
    //  container.appendChild(draggable);
    //} else {
    //  container.insertBefore(draggable, afterElement);
    //}
  });
});

function getDraggableElement(container, y) {
  const draggableElements = [
    ...container.querySelectorAll(".draggable:not(.dragging)"),
  ];

  return draggableElements.reduce(
    (closest, child) => {
      const box = child.getBoundingClientRect();
      const offset = y - box.top - box.height / 2;
      if (offset < 0 && offset > closest.offset) {
        return { offset, element: child };
      } else {
        return closest;
      }
    },
    {
      offset: Number.NEGATIVE_INFINITY,
    },
  ).element;
}
