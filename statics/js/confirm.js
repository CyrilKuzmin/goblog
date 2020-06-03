//Функция для подтверждения желания юзера удалить пост

function confirmRemoval() {
    if (confirm("Вы уверены, что хотите удалить этот пост?")) {
        return true;
    } else {
        return false;
    }
}