/**
 * Общие функции для сайта
 */

// Инициализация страницы
document.addEventListener('DOMContentLoaded', function() {
    // Плавное появление контента
    document.querySelector('.main').classList.add('fade-in');
    
    // Выделение активной ссылки в навигации
    const currentPath = window.location.pathname;
    document.querySelectorAll('.header__link').forEach(link => {
        if (link.getAttribute('href') === currentPath) {
            link.classList.add('header__link--active');
        }
    });
    
    // Управление отображением сообщений
    const urlParams = new URLSearchParams(window.location.search);
    const successMsg = urlParams.get('success');
    const errorMsg = urlParams.get('error');
    
    if (successMsg) {
        showMessage(successMsg, 'success');
    }
    
    if (errorMsg) {
        showMessage(errorMsg, 'error');
    }
});

/**
 * Показать сообщение
 * @param {string} text - Текст сообщения
 * @param {string} type - Тип (success, error)
 */
function showMessage(text, type) {
    const messageDiv = document.createElement('div');
    messageDiv.className = `global-message ${type}-message`;
    messageDiv.textContent = text;
    
    document.body.prepend(messageDiv);
    
    // Автоматическое скрытие через 5 секунд
    setTimeout(() => {
        messageDiv.style.opacity = '0';
        setTimeout(() => messageDiv.remove(), 300);
    }, 5000);
}

/**
 * Проверка авторизации
 * @returns {boolean} Авторизован ли пользователь
 */
function isAuthenticated() {
    return localStorage.getItem('token') !== null;
}

/**
 * Перенаправление на страницу входа при отсутствии авторизации
 */
function requireAuth() {
    if (!isAuthenticated()) {
        window.location.href = '/cgi-bin/login.cgi?error=Требуется авторизация';
    }
}

/**
 * Выход из системы
 */
function logout() {
    localStorage.removeItem('token');
    window.location.href = '/cgi-bin/index.cgi';
}