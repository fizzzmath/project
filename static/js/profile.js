/**
 * Обработка профиля пользователя
 */

document.addEventListener('DOMContentLoaded', function() {
    // Проверка авторизации для страниц профиля
    if (window.location.pathname.includes('/profile')) {
        requireAuth();
    }
    
    // Обработчик формы редактирования профиля
    const profileForm = document.getElementById('profileForm');
    if (profileForm) {
        profileForm.addEventListener('submit', handleProfileUpdate);
    }
});

/**
 * Обновление профиля
 * @param {Event} e - Событие формы
 */
function handleProfileUpdate(e) {
    e.preventDefault();
    
    const form = e.target;
    const formData = new FormData(form);
    const data = {
        fullName: formData.get('fullName'),
        email: formData.get('email'),
        bio: formData.get('bio')
    };
    
    const userId = form.dataset.userId;
    
    fetch(`/binaries/profile.cgi?user_id=${userId}&action=edit`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
        body: JSON.stringify(data)
    })
    // .then(response => {
    //     if (!response.ok) {
    //         return response.json().then(err => { throw err; });
    //     }
    //     return response.json();
    // })
    // .then(data => {
    //     showFormSuccess(form, 'Профиль успешно обновлен');
    // })
    // .catch(error => {
    //     showFormError(form, error.error || error.message || 'Ошибка обновления профиля');
    // });
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
        
        // Скрыть сообщение об успехе, если есть
        const successDiv = form.querySelector('.success-message');
        if (successDiv) successDiv.style.display = 'none';
    }
}

/**
 * Показать успешное сообщение
 * @param {HTMLFormElement} form - Форма
 * @param {string} message - Сообщение
 */
function showFormSuccess(form, message) {
    const successDiv = form.querySelector('.success-message');
    if (successDiv) {
        successDiv.textContent = message;
        successDiv.style.display = 'block';
        
        // Скрыть сообщение об ошибке, если есть
        const errorDiv = form.querySelector('.error-message');
        if (errorDiv) errorDiv.style.display = 'none';
    }
}