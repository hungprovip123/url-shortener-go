// URL shortener frontend logic

const API_BASE = '/api/v1';

// DOM elements
const urlInput = document.getElementById('urlInput');
const shortenBtn = document.getElementById('shortenBtn');
const errorDiv = document.getElementById('error');
const resultDiv = document.getElementById('result');
const urlsList = document.getElementById('urlsList');

// Initialize when page loads
document.addEventListener('DOMContentLoaded', function() {
    loadUrls();
    
    // Enter key support
    urlInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            shortenUrl();
        }
    });
});

// Shorten URL function
async function shortenUrl() {
    const url = urlInput.value.trim();
    
    if (!url) {
        showError('Vui lòng nhập URL cần rút gọn');
        return;
    }
    
    if (!isValidUrl(url)) {
        showError('URL không hợp lệ. Vui lòng nhập URL đầy đủ (bao gồm http:// hoặc https://)');
        return;
    }
    
    hideMessages();
    setLoading(true);
    
    try {
        const response = await fetch(`${API_BASE}/shorten`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url: url })
        });
        
        const data = await response.json();
        
        if (response.ok) {
            showResult(data);
            urlInput.value = '';
            loadUrls(); // Refresh the list
        } else {
            showError(data.error || 'Có lỗi xảy ra khi rút gọn URL');
        }
    } catch (error) {
        showError('Không thể kết nối đến server. Vui lòng thử lại sau.');
        console.error('Error:', error);
    } finally {
        setLoading(false);
    }
}

// Load URLs list
async function loadUrls() {
    try {
        urlsList.innerHTML = '<div class="loading">Đang tải...</div>';
        
        const response = await fetch(`${API_BASE}/urls`);
        const data = await response.json();
        
        if (response.ok) {
            displayUrls(data);
        } else {
            urlsList.innerHTML = '<div class="error">Không thể tải danh sách URL</div>';
        }
    } catch (error) {
        urlsList.innerHTML = '<div class="error">Lỗi kết nối server</div>';
        console.error('Error loading URLs:', error);
    }
}

// Display URLs in the list
function displayUrls(urls) {
    if (!urls || urls.length === 0) {
        urlsList.innerHTML = `
            <div class="empty-state">
                <h3>📝 Chưa có URL nào</h3>
                <p>Hãy tạo URL ngắn đầu tiên của bạn!</p>
            </div>
        `;
        return;
    }
    
    const urlsHtml = urls.map(url => `
        <div class="url-item">
            <div class="url-header">
                <div class="url-info">
                    <h3>🔗 ${truncateUrl(url.original_url, 50)}</h3>
                    <p>URL ngắn: <span class="short-url">${url.short_url}</span></p>
                    <p><small>📅 Tạo lúc: ${formatDate(url.created_at)}</small></p>
                </div>
                <div class="url-stats">
                    👆 ${url.clicks} lượt click
                </div>
            </div>
            <div class="url-actions">
                <button class="btn btn-primary" onclick="copyToClipboard('${url.short_url}')">
                    📋 Copy
                </button>
                <a href="${url.short_url}" target="_blank" class="btn btn-secondary">
                    🚀 Mở
                </a>
                <button class="btn btn-secondary" onclick="getStats('${url.short_code}')">
                    📊 Thống kê
                </button>
            </div>
        </div>
    `).join('');
    
    urlsList.innerHTML = urlsHtml;
}

// Get URL statistics
async function getStats(shortCode) {
    try {
        const response = await fetch(`${API_BASE}/stats/${shortCode}`);
        const data = await response.json();
        
        if (response.ok) {
            alert(`📊 Thống kê cho ${data.short_code}:\n\n` +
                  `🔗 URL gốc: ${data.original_url}\n` +
                  `📱 URL ngắn: ${data.short_url}\n` +
                  `👆 Số lượt click: ${data.clicks}\n` +
                  `📅 Tạo lúc: ${formatDate(data.created_at)}`);
        } else {
            alert('❌ Không thể lấy thống kê cho URL này');
        }
    } catch (error) {
        alert('❌ Lỗi kết nối server');
        console.error('Error getting stats:', error);
    }
}

// Copy to clipboard
async function copyToClipboard(text) {
    try {
        await navigator.clipboard.writeText(text);
        showResult({ short_url: text, message: 'Đã copy vào clipboard!' });
    } catch (error) {
        // Fallback for older browsers
        const textArea = document.createElement('textarea');
        textArea.value = text;
        document.body.appendChild(textArea);
        textArea.select();
        document.execCommand('copy');
        document.body.removeChild(textArea);
        showResult({ short_url: text, message: 'Đã copy vào clipboard!' });
    }
}

// Utility functions
function isValidUrl(string) {
    try {
        new URL(string);
        return true;
    } catch (_) {
        return false;
    }
}

function truncateUrl(url, maxLength) {
    if (url.length <= maxLength) return url;
    return url.substring(0, maxLength) + '...';
}

function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleString('vi-VN', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
}

function showError(message) {
    hideMessages();
    errorDiv.textContent = '❌ ' + message;
    errorDiv.classList.remove('hidden');
}

function showResult(data) {
    hideMessages();
    resultDiv.innerHTML = `
        <div style="margin-bottom: 1rem;">
            <strong>✅ ${data.message || 'URL đã được rút gọn thành công!'}</strong>
        </div>
        <div style="margin-bottom: 1rem;">
            <strong>URL ngắn:</strong> 
            <span class="short-url">${data.short_url}</span>
        </div>
        <div style="text-align: center;">
            <button class="btn btn-primary" onclick="copyToClipboard('${data.short_url}')">
                📋 Copy URL
            </button>
        </div>
    `;
    resultDiv.classList.remove('hidden');
}

function hideMessages() {
    errorDiv.classList.add('hidden');
    resultDiv.classList.add('hidden');
}

function setLoading(loading) {
    shortenBtn.disabled = loading;
    shortenBtn.textContent = loading ? '⏳ Đang xử lý...' : 'Rút gọn';
} 