/**
 * Обработка форм авторизации
 */

document.addEventListener('DOMContentLoaded', function() {
    // Обработчик формы входа
    const loginForm = document.getElementById('loginForm');
    if (loginForm) {
        loginForm.addEventListener('submit', handleLogin);
    }
    
    // Обработчик формы регистрации
    const registerForm = document.getElementById('registerForm');
    if (registerForm) {
        registerForm.addEventListener('submit', handleRegister);
    }
});

/**
 * Обработка входа
 * @param {Event} e - Событие формы
 */
function handleLogin(e) {
    e.preventDefault();
    
    const form = e.target;
    const formData = new FormData(form);
    const data = {
        username: formData.get('username'),
        password: formData.get('password')
    };
    
    fetch('/binaries/login.cgi', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(err => { throw err; });
        }
        return response.json();
    })
    .then(data => {
        // Сохраняем токен
        localStorage.setItem('token', data.token);
        
        // Перенаправляем на главную
        window.location.href = '/binaries/index.cgi';
    })
    .catch(error => {
        showFormError(form, error.message || 'Ошибка входа');
    });
}

/**
 * Обработка регистрации
 * @param {Event} e - Событие формы
 */
function handleRegister(e) {
    e.preventDefault();
    
    const form = e.target;
    const formData = new FormData(form);
    const data = {
        username: formData.get('username'),
        email: formData.get('email'),
        password: formData.get('password'),
        confirmPassword: formData.get('confirmPassword')
    };
    
    // Базовая проверка паролей
    if (data.password !== data.confirmPassword) {
        showFormError(form, 'Пароли не совпадают');
        return;
    }
    
    fetch('/binaries/register.cgi', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username: data.username,
            email: data.email,
            password: data.password
        })
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(err => { throw err; });
        }
        return response.json();
    })
    .then(data => {
        // Перенаправляем на страницу входа
        window.location.href = '/templates/login.cgi?success=Регистрация прошла успешно';
    })
    .catch(error => {
        showFormError(form, error.message || 'Ошибка регистрации');
    });
}

/**
 * Показать ошибку формы
 * @param {HTMLFormElement} form - Форма
 * @param {string} message - Сообщение об ошибке
 */
function showFormError(form, message) {
    const errorDiv = form.querySelector('.error-message');
    if (errorDiv) {
        errorDiv.textContent = message;
        errorDiv.style.display = 'block';
    }
}