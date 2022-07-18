const formEl = document.querySelector(".form");
const inputEl = document.querySelector(".input");
const ulEL = document.querySelector(".list");

let list = JSON.parse(localStorage.getItem("list"));
// console.log(list);
if(list){
    list.forEach((task) => {
        toDoList(task);
    });
}

formEl.addEventListener("submit", (event) => {
    //prevents the page to be refreshed
    event.preventDefault();
    // console.log(inputEl.value);
    toDoList();

});

function toDoList(task){
    let newTask = inputEl.value;
    if (task){
        newTask = task.name;
    }

    const liEl= document.createElement("li");
    if (task && task.checked){
        liEl.classList.add("checked");
        
    }
    liEl.innerText = newTask;
    ulEL.appendChild(liEl);
    inputEl.value = "";

    const checkBtnEl = document.createElement("div");
    checkBtnEl.innerHTML = `
    <i class="fas fa-check-square"></i>
    `;
    liEl.appendChild(checkBtnEl);

    const trashBtnEl = document.createElement("div");
    trashBtnEl.innerHTML = `
    <i class="fas fa-trash"></i></li>
    `;
    liEl.appendChild(trashBtnEl);

    checkBtnEl.addEventListener("click", () => {
        liEl.classList.toggle("checked");
        updateLocalStorage();
    });

    trashBtnEl.addEventListener("click", () => {
        liEl.remove();
        updateLocalStorage();
    });
    updateLocalStorage();
};

function updateLocalStorage(){
    const liEls = document.querySelectorAll("li");
    list = [];
    liEls.forEach((liEl) => {
        list.push({
            name: liEl.innerText,
            checked: liEl.classList.contains("checked")
        });
    });
    localStorage.setItem("list", JSON.stringify(list));

}