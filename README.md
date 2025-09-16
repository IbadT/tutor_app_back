# Tutor App - Backend API Documentation

## Архитектура приложения

Приложение построено на основе **Clean Architecture** принципов с четким разделением на слои:

### Структура проекта

```
internal/
├── app/                    # Application Layer (Приложение)
│   ├── handlers/          # HTTP обработчики запросов
│   │   ├── auth.go       # Обработчики аутентификации
│   │   └── user.go       # Обработчики пользователей
│   ├── middleware/        # HTTP middleware
│   │   ├── auth.go       # Аутентификация и авторизация
│   │   └── error_handler.go # Обработка ошибок
│   └── server.go         # Настройка и запуск сервера
├── domain/               # Domain Layer (Доменная логика)
│   ├── auth/            # Доменная логика аутентификации
│   │   ├── service.go   # Бизнес-логика аутентификации
│   │   ├── repository.go # Интерфейсы репозиториев
│   │   └── types.go     # Доменные типы и структуры
│   ├── user/            # Доменная логика пользователей
│   │   ├── service.go   # Бизнес-логика пользователей
│   │   ├── repository.go # Интерфейсы репозиториев
│   │   └── types.go     # Доменные типы и структуры
│   └── shared/          # Общие доменные компоненты
│       └── errors.go    # Доменные ошибки
├── infrastructure/       # Infrastructure Layer (Инфраструктура)
│   ├── database/        # Подключение к базе данных
│   │   └── connection.go
│   ├── repositories/    # Реализация репозиториев
│   │   ├── auth.go      # Репозиторий аутентификации
│   │   └── user.go      # Репозиторий пользователей
│   └── external/        # Внешние сервисы
│       ├── jwt.go       # JWT токены
│       └── password.go  # Хеширование паролей
└── config/              # Конфигурация приложения
    └── config.go
```

### Описание слоев

#### 1. Domain Layer (Доменный слой)
**Назначение**: Содержит бизнес-логику и правила приложения
- **types.go**: Доменные сущности, DTO, интерфейсы
- **service.go**: Бизнес-логика, валидация, обработка данных
- **repository.go**: Интерфейсы для доступа к данным
- **shared/errors.go**: Доменные ошибки и их обработка

#### 2. Infrastructure Layer (Инфраструктурный слой)
**Назначение**: Техническая реализация внешних зависимостей
- **database/**: Подключение к БД, миграции
- **repositories/**: Реализация интерфейсов репозиториев (GORM)
- **external/**: Внешние сервисы (JWT, хеширование, API)

#### 3. Application Layer (Слой приложения)
**Назначение**: Координация между слоями, HTTP обработка
- **handlers/**: HTTP обработчики, валидация запросов
- **middleware/**: Аутентификация, логирование, CORS
- **server.go**: Инициализация сервера, DI контейнер

### Принципы архитектуры

1. **Dependency Inversion**: Зависимости направлены внутрь
2. **Interface Segregation**: Маленькие, специфичные интерфейсы
3. **Single Responsibility**: Каждый пакет отвечает за одну область
4. **Clean Architecture**: Четкое разделение слоев
5. **Dependency Injection**: Внедрение зависимостей через конструкторы

### Поток данных

```
HTTP Request → Handler → Service → Repository → Database
                ↓         ↓         ↓
HTTP Response ← Handler ← Service ← Repository ← Database
```

## Обзор приложения

Это образовательная платформа для студентов и преподавателей с функциями:
- Управление курсами и материалами
- Форумы для обсуждений
- Отслеживание прогресса обучения
- Система достижений и рейтингов
- Профили пользователей
- Уведомления

## Структура API

### Базовый URL
```
https://api.tutor-app.com/v1
```

### Аутентификация
Все запросы (кроме публичных) требуют JWT токен в заголовке:
```
Authorization: Bearer <jwt_token>
```










---

❎ ## 1. Аутентификация и пользователи

### 1.1 Регистрация
```
POST /auth/register
```

**Тело запроса:**
```json
{
  "firstName": "Алексей",
  "lastName": "Петров",
  "email": "alexey.petrov@example.com",
  "password": "SecurePassword123!",
  "role": "student",
  "location": "Москва, Россия"
}
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "user_123",
      "firstName": "Алексей",
      "lastName": "Петров",
      "email": "alexey.petrov@example.com",
      "role": "student",
      "isVerified": false,
      "joinedAt": "2024-01-15T10:30:00Z"
    },
    "tokens": {
      "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
}
```

❎ ### 1.2 Вход
```
POST /auth/login
```

**Тело запроса:**
```json
{
  "email": "alexey.petrov@example.com",
  "password": "SecurePassword123!"
}
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "user_123",
      "firstName": "Алексей",
      "lastName": "Петров",
      "email": "alexey.petrov@example.com",
      "role": "student",
      "avatar": "https://api.tutor-app.com/avatars/user_123.jpg",
      "isVerified": true
    },
    "tokens": {
      "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
}
```


❌
❎





❌ ### 1.3 Восстановление пароля
```
POST /auth/forgot-password
```

**Тело запроса:**
```json
{
  "email": "alexey.petrov@example.com"
}
```






❎ ### 1.4 Обновление токена
```
POST /auth/refresh
```

**Тело запроса:**
```json
{
  "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

❎ ## 2. Профиль пользователя

### 2.1 Получить профиль
```
GET /users/profile
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "id": "user_123",
    "name": "Алексей Петров",
    "email": "alexey.petrov@example.com",
    "role": "student",
    "avatar": "https://api.tutor-app.com/avatars/user_123.jpg",
    "bio": "Студент факультета информатики",
    "location": "Москва, Россия",
    "joinedAt": "2023-09-01T00:00:00Z",
    "coursesCompleted": 12,
    "coursesInProgress": 3,
    "achievements": 8,
    "followers": 156,
    "following": 89,
    "level": 15,
    "xp": 2340,
    "nextLevelXp": 3000,
    "isVerified": true,
    "badges": ["Первые шаги", "Серьезный студент", "Помощник сообщества"]
  }
}
```

❎ ### 2.2 Обновить профиль
```
PUT /users/profile
```

**Тело запроса:**
```json
{
  "firstName": "Алексей",
  "lastName": "Петров",
  "bio": "Обновленная биография",
  "location": "Санкт-Петербург, Россия",
  "phone": "+7 (999) 123-45-67"
}
```








### 2.3 Загрузить аватар
```
POST /users/profile/avatar
```

**Форма данных:**
- `avatar`: файл изображения (multipart/form-data)

### 2.4 Статистика профиля
```
GET /users/profile/stats
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "studyTime": {
      "total": "127 часов",
      "thisWeek": "8.5 часов",
      "thisMonth": "34 часа",
      "averagePerDay": "2.1 часа"
    },
    "achievements": {
      "total": 8,
      "thisMonth": 2,
      "recent": ["Завершил курс React", "Помог 5 студентам"]
    },
    "courses": {
      "completed": 12,
      "inProgress": 3,
      "enrolled": 15,
      "averageRating": 4.7
    },
    "activity": {
      "streak": 7,
      "longestStreak": 21,
      "weeklyGoal": 10,
      "weeklyProgress": 8.5
    },
    "performance": {
      "averageScore": 87,
      "assignmentsCompleted": 45,
      "quizzesPassed": 23,
      "projectsSubmitted": 8
    }
  }
}
```

### 2.5 Детальная статистика
```
GET /users/profile/detailed-stats
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "weeklyProgress": {
      "labels": ["Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"],
      "data": [2, 3, 1.5, 4, 2.5, 1, 3]
    },
    "monthlyProgress": {
      "labels": ["Янв", "Фев", "Мар", "Апр", "Май", "Июн"],
      "data": [45, 52, 48, 61, 55, 67]
    },
    "courseDistribution": {
      "completed": 12,
      "inProgress": 3,
      "notStarted": 5
    },
    "studyTimeDistribution": {
      "videos": 40,
      "reading": 30,
      "practice": 20,
      "quizzes": 10
    },
    "achievementsTimeline": [
      {
        "date": "2024-01-15",
        "achievement": "Завершил курс React",
        "points": 100
      }
    ],
    "performanceMetrics": {
      "averageScore": 87,
      "improvement": 12,
      "consistency": 85,
      "engagement": 92
    },
    "studyStreaks": {
      "current": 7,
      "longest": 21,
      "average": 5.2,
      "totalDays": 156
    }
  }
}
```

---

## 3. Настройки профиля

### 3.1 Получить настройки профиля
```
GET /users/profile/settings
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "personalInfo": {
      "firstName": "Алексей",
      "lastName": "Петров",
      "email": "alexey.petrov@example.com",
      "phone": "+7 (999) 123-45-67",
      "location": "Москва, Россия",
      "bio": "Студент факультета информатики",
      "dateOfBirth": "1998-05-15"
    },
    "account": {
      "username": "alexey_petrov"
    },
    "preferences": {
      "language": "ru",
      "timezone": "Europe/Moscow",
      "emailNotifications": true,
      "smsNotifications": false,
      "publicProfile": true,
      "showEmail": false,
      "showLocation": true
    }
  }
}
```

### 3.2 Обновить настройки профиля
```
PUT /users/profile/settings
```

**Тело запроса:**
```json
{
  "personalInfo": {
    "firstName": "Алексей",
    "lastName": "Петров",
    "phone": "+7 (999) 123-45-67",
    "location": "Санкт-Петербург, Россия",
    "bio": "Обновленная биография"
  },
  "preferences": {
    "language": "en",
    "timezone": "Europe/London",
    "emailNotifications": true,
    "smsNotifications": true,
    "publicProfile": false
  }
}
```







❎ ### 3.3 Сменить пароль
```
PUT /users/profile/password
```

**Тело запроса:**
```json
{
  "currentPassword": "OldPassword123!",
  "newPassword": "NewPassword123!"
}
```

---

## 4. Настройки уведомлений

### 4.1 Получить настройки уведомлений
```
GET /users/profile/notifications
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "general": {
      "enableAll": true,
      "sound": true,
      "desktop": true,
      "mobile": true
    },
    "courseUpdates": {
      "newLessons": true,
      "assignments": true,
      "deadlines": true,
      "grades": true,
      "announcements": true
    },
    "social": {
      "messages": true,
      "friendRequests": true,
      "comments": false,
      "mentions": true,
      "follows": false
    },
    "achievements": {
      "badges": true,
      "levels": true,
      "streaks": true,
      "milestones": true
    },
    "system": {
      "maintenance": true,
      "updates": true,
      "security": true,
      "newsletters": false
    },
    "schedule": {
      "quietHours": false,
      "quietStart": "22:00",
      "quietEnd": "08:00",
      "weekdaysOnly": false
    }
  }
}
```

### 4.2 Обновить настройки уведомлений
```
PUT /users/profile/notifications
```

**Тело запроса:**
```json
{
  "general": {
    "enableAll": true,
    "sound": false,
    "desktop": true,
    "mobile": false
  },
  "courseUpdates": {
    "newLessons": true,
    "assignments": false,
    "deadlines": true,
    "grades": true,
    "announcements": false
  },
  "schedule": {
    "quietHours": true,
    "quietStart": "23:00",
    "quietEnd": "07:00",
    "weekdaysOnly": true
  }
}
```

---

## 5. Настройки приватности

### 5.1 Получить настройки приватности
```
GET /users/profile/privacy
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "profileVisibility": {
      "publicProfile": true,
      "showEmail": false,
      "showPhone": false,
      "showLocation": true,
      "showBirthDate": false,
      "showBio": true
    },
    "activityPrivacy": {
      "showOnlineStatus": true,
      "showLastSeen": true,
      "showStudyActivity": true,
      "showAchievements": true,
      "showCourseProgress": false
    },
    "socialPrivacy": {
      "allowFriendRequests": true,
      "allowMessages": true,
      "allowComments": true,
      "allowMentions": true,
      "showFollowers": true,
      "showFollowing": true
    },
    "dataSharing": {
      "analytics": true,
      "marketing": false,
      "research": false,
      "thirdParty": false,
      "locationData": false
    },
    "accountSecurity": {
      "twoFactorAuth": false,
      "loginAlerts": true,
      "sessionTimeout": "30",
      "passwordExpiry": "90"
    }
  }
}
```

### 5.2 Обновить настройки приватности
```
PUT /users/profile/privacy
```

**Тело запроса:**
```json
{
  "profileVisibility": {
    "publicProfile": false,
    "showEmail": false,
    "showLocation": false
  },
  "socialPrivacy": {
    "allowFriendRequests": false,
    "allowMessages": false
  },
  "dataSharing": {
    "analytics": false,
    "marketing": false
  }
}
```













---

## 6. Курсы

### 6.1 Получить список курсов пользователя
```
GET /courses
```

**Параметры запроса:**
- `search` (string, optional): Поиск по названию
- `category` (string, optional): Фильтр по категории
- `status` (string, optional): `all`, `enrolled`, `completed`, `in_progress`
- `sort` (string, optional): `newest`, `oldest`, `name`, `progress`
- `page` (number, optional): Номер страницы (по умолчанию 1)
- `limit` (number, optional): Количество на странице (по умолчанию 10)

**Ответ:**
```json
{
  "success": true,
  "data": {
    "courses": [
      {
        "id": "course_123",
        "title": "React для начинающих",
        "instructor": "Др. Анна Иванова",
        "description": "Полный курс изучения React с нуля",
        "progress": 75,
        "totalLessons": 24,
        "completedLessons": 18,
        "duration": "4 недели",
        "studentsCount": 1250,
        "rating": 4.8,
        "nextLesson": {
          "title": "Хуки в React",
          "time": "2024-01-20T14:00:00Z"
        },
        "category": "informatics",
        "thumbnail": "https://api.tutor-app.com/courses/course_123/thumbnail.jpg"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 25,
      "pages": 3
    }
  }
}
```

### 6.2 Получить детали курса
```
GET /courses/:courseId
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "id": "course_123",
    "title": "React для начинающих",
    "instructor": {
      "id": "teacher_456",
      "name": "Др. Анна Иванова",
      "avatar": "https://api.tutor-app.com/avatars/teacher_456.jpg"
    },
    "description": "Полный курс изучения React с нуля",
    "progress": 75,
    "totalLessons": 24,
    "completedLessons": 18,
    "duration": "4 недели",
    "studentsCount": 1250,
    "rating": 4.8,
    "category": "informatics",
    "thumbnail": "https://api.tutor-app.com/courses/course_123/thumbnail.jpg",
    "lessons": [
      {
        "id": "lesson_1",
        "title": "Введение в React",
        "duration": "45 минут",
        "isCompleted": true,
        "order": 1
      },
      {
        "id": "lesson_2",
        "title": "Компоненты",
        "duration": "60 минут",
        "isCompleted": true,
        "order": 2
      }
    ],
    "assignments": [
      {
        "id": "assignment_1",
        "title": "Создание первого компонента",
        "dueDate": "2024-01-25T23:59:59Z",
        "isSubmitted": true,
        "grade": 95
      }
    ]
  }
}
```

### 6.3 Записаться на курс
```
POST /courses/:courseId/enroll
```

**Ответ:**
```json
{
  "success": true,
  "message": "Успешно записались на курс"
}
```

### 6.4 Отметить урок как завершенный
```
POST /courses/:courseId/lessons/:lessonId/complete
```

### 6.5 Получить статистику курсов
```
GET /courses/stats
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "totalCourses": 25,
    "completedCourses": 12,
    "inProgressCourses": 3,
    "averageProgress": 68.5,
    "totalStudyTime": "127 часов",
    "thisWeekStudyTime": "8.5 часов"
  }
}
```

---

## 7. Учебные материалы

### 7.1 Получить список материалов
```
GET /materials
```

**Параметры запроса:**
- `search` (string, optional): Поиск по названию или описанию
- `type` (string, optional): `all`, `video`, `document`, `image`, `audio`, `archive`, `link`, `book`
- `category` (string, optional): `all`, `programming`, `mathematics`, `physics`, `chemistry`, `biology`, `general`
- `sort` (string, optional): `newest`, `oldest`, `popular`, `name`, `size`
- `page` (number, optional): Номер страницы
- `limit` (number, optional): Количество на странице

**Ответ:**
```json
{
  "success": true,
  "data": {
    "materials": [
      {
        "id": "material_123",
        "title": "React Hooks Complete Guide",
        "description": "Comprehensive guide covering all React hooks",
        "type": "document",
        "category": "programming",
        "size": "2.4 MB",
        "duration": null,
        "author": {
          "name": "Др. Сара Джонсон",
          "role": "teacher"
        },
        "tags": ["react", "hooks", "javascript", "frontend"],
        "downloadCount": 1247,
        "rating": 4.8,
        "isDownloaded": false,
        "isBookmarked": true,
        "uploadedAt": "2024-01-15T10:30:00Z",
        "url": "https://api.tutor-app.com/materials/material_123/download"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 156,
      "pages": 16
    }
  }
}
```

### 7.2 Загрузить материал
```
POST /materials
```

**Форма данных:**
- `title` (string): Название материала
- `description` (string): Описание
- `type` (string): Тип материала
- `category` (string): Категория
- `tags` (string): Теги через запятую
- `file` (file): Файл материала (multipart/form-data)

### 7.3 Скачать материал
```
GET /materials/:materialId/download
```

### 7.4 Добавить материал в закладки
```
POST /materials/:materialId/bookmark
```

### 7.5 Оценить материал
```
POST /materials/:materialId/rate
```

**Тело запроса:**
```json
{
  "rating": 5,
  "review": "Отличный материал!"
}
```

### 7.6 Получить статистику материалов
```
GET /materials/stats
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "totalMaterials": 156,
    "totalDownloads": 15420,
    "totalSize": "512 MB",
    "averageRating": 4.6
  }
}
```

---

## 8. Форумы и обсуждения

### 8.1 Получить список обсуждений
```
GET /discussions
```

**Параметры запроса:**
- `search` (string, optional): Поиск по названию или содержимому
- `category` (string, optional): `all`, `general`, `questions`, `announcements`, `feedback`, `projects`, `resources`
- `sort` (string, optional): `newest`, `oldest`, `popular`, `trending`, `replies`
- `page` (number, optional): Номер страницы
- `limit` (number, optional): Количество на странице

**Ответ:**
```json
{
  "success": true,
  "data": {
    "discussions": [
      {
        "id": "discussion_123",
        "title": "Как решить задачу по алгоритмам?",
        "content": "Помогите разобраться с задачей...",
        "author": {
          "name": "Алексей Петров",
          "avatar": "https://api.tutor-app.com/avatars/user_123.jpg",
          "role": "student"
        },
        "category": "questions",
        "tags": ["алгоритмы", "программирование", "помощь"],
        "repliesCount": 5,
        "viewsCount": 23,
        "likesCount": 8,
        "isLiked": false,
        "isBookmarked": true,
        "createdAt": "2024-01-15T14:30:00Z",
        "lastActivity": "2024-01-16T09:15:00Z",
        "isPinned": false,
        "isSolved": false
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 89,
      "pages": 9
    }
  }
}
```

### 8.2 Создать новое обсуждение
```
POST /discussions
```

**Тело запроса:**
```json
{
  "title": "Как решить задачу по алгоритмам?",
  "content": "Помогите разобраться с задачей сортировки...",
  "category": "questions",
  "tags": ["алгоритмы", "программирование", "помощь"]
}
```

### 8.3 Получить детали обсуждения
```
GET /discussions/:discussionId
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "id": "discussion_123",
    "title": "Как решить задачу по алгоритмам?",
    "content": "Помогите разобраться с задачей...",
    "author": {
      "name": "Алексей Петров",
      "avatar": "https://api.tutor-app.com/avatars/user_123.jpg",
      "role": "student"
    },
    "category": "questions",
    "tags": ["алгоритмы", "программирование", "помощь"],
    "repliesCount": 5,
    "viewsCount": 23,
    "likesCount": 8,
    "isLiked": false,
    "isBookmarked": true,
    "createdAt": "2024-01-15T14:30:00Z",
    "lastActivity": "2024-01-16T09:15:00Z",
    "replies": [
      {
        "id": "reply_1",
        "content": "Попробуйте использовать быструю сортировку...",
        "author": {
          "name": "Др. Анна Иванова",
          "avatar": "https://api.tutor-app.com/avatars/teacher_456.jpg",
          "role": "teacher"
        },
        "createdAt": "2024-01-15T16:45:00Z",
        "likesCount": 3,
        "isLiked": false
      }
    ]
  }
}
```

### 8.4 Ответить на обсуждение
```
POST /discussions/:discussionId/replies
```

**Тело запроса:**
```json
{
  "content": "Попробуйте использовать быструю сортировку..."
}
```

### 8.5 Лайкнуть обсуждение
```
POST /discussions/:discussionId/like
```

### 8.6 Добавить обсуждение в закладки
```
POST /discussions/:discussionId/bookmark
```

### 8.7 Отметить обсуждение как решенное
```
POST /discussions/:discussionId/solve
```

---

## 9. Отслеживание прогресса

### 9.1 Получить прогресс обучения
```
GET /progress
```

**Ответ:**
```json
{
  "success": true,
  "data": {
    "overview": {
      "totalStudyTime": "127 часов",
      "thisWeekStudyTime": "8.5 часов",
      "coursesCompleted": 12,
      "coursesInProgress": 3,
      "averageScore": 87,
      "currentStreak": 7
    },
    "weeklyProgress": [
      {
        "date": "2024-01-15",
        "studyTime": 2.5,
        "lessonsCompleted": 3,
        "assignmentsSubmitted": 1
      }
    ],
    "achievements": [
      {
        "id": "achievement_1",
        "name": "Первые шаги",
        "description": "Завершили первый урок",
        "earnedAt": "2024-01-10T10:00:00Z",
        "icon": "🎯"
      }
    ],
    "goals": [
      {
        "id": "goal_1",
        "title": "Изучить React",
        "description": "Завершить курс React для начинающих",
        "targetDate": "2024-02-15",
        "progress": 75,
        "isCompleted": false
      }
    ]
  }
}
```

### 9.2 Обновить цель обучения
```
PUT /progress/goals/:goalId
```

**Тело запроса:**
```json
{
  "title": "Изучить React",
  "description": "Завершить курс React для начинающих",
  "targetDate": "2024-02-15"
}
```

### 9.3 Создать новую цель
```
POST /progress/goals
```

**Тело запроса:**
```json
{
  "title": "Изучить TypeScript",
  "description": "Освоить основы TypeScript",
  "targetDate": "2024-03-01"
}
```

---

## 10. Уведомления

### 10.1 Получить уведомления
```
GET /notifications
```

**Параметры запроса:**
- `page` (number, optional): Номер страницы
- `limit` (number, optional): Количество на странице
- `unread` (boolean, optional): Только непрочитанные

**Ответ:**
```json
{
  "success": true,
  "data": {
    "notifications": [
      {
        "id": "notification_123",
        "title": "Новый урок доступен",
        "body": "В курсе 'React для начинающих' появился новый урок",
        "type": "info",
        "timestamp": "2024-01-15T10:30:00Z",
        "read": false,
        "actionUrl": "/courses/course_123/lessons/lesson_5"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 25,
      "pages": 3
    }
  }
}
```

### 10.2 Отметить уведомление как прочитанное
```
PUT /notifications/:notificationId/read
```

### 10.3 Отметить все уведомления как прочитанные
```
PUT /notifications/read-all
```

### 10.4 Удалить уведомление
```
DELETE /notifications/:notificationId
```

---

## 11. Файлы и загрузки

### 11.1 Загрузить файл
```
POST /files/upload
```

**Форма данных:**
- `file` (file): Файл для загрузки (multipart/form-data)
- `type` (string): Тип файла (`material`, `avatar`, `course_thumbnail`)

**Ответ:**
```json
{
  "success": true,
  "data": {
    "id": "file_123",
    "url": "https://api.tutor-app.com/files/file_123",
    "filename": "document.pdf",
    "size": 2048576,
    "mimeType": "application/pdf"
  }
}
```

### 11.2 Получить файл
```
GET /files/:fileId
```

---

## 12. Система достижений

### 12.1 Получить достижения пользователя
```
GET /achievements
```

**Ответ:**
```json
{
  "success": true,
  "data": [
    {
      "id": "achievement_1",
      "name": "Первые шаги",
      "description": "Завершили первый урок",
      "icon": "🎯",
      "earnedAt": "2024-01-10T10:00:00Z",
      "points": 10
    },
    {
      "id": "achievement_2",
      "name": "Серьезный студент",
      "description": "Завершили 5 курсов",
      "icon": "🎓",
      "earnedAt": "2024-01-12T15:30:00Z",
      "points": 50
    }
  ]
}
```

### 12.2 Получить доступные достижения
```
GET /achievements/available
```

---

## Коды ошибок

### HTTP статус коды
- `200` - Успешный запрос
- `201` - Ресурс создан
- `400` - Неверный запрос
- `401` - Не авторизован
- `403` - Доступ запрещен
- `404` - Ресурс не найден
- `409` - Конфликт (например, email уже существует)
- `422` - Ошибка валидации
- `500` - Внутренняя ошибка сервера

### Формат ошибок
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Ошибка валидации данных",
    "details": {
      "email": "Email уже используется",
      "password": "Пароль должен содержать минимум 8 символов"
    }
  }
}
```

---

## Лимиты и ограничения

### Rate Limiting
- Общие запросы: 1000 запросов в час
- Загрузка файлов: 100 запросов в час
- Регистрация/вход: 10 запросов в час

### Размеры файлов
- Аватары: максимум 5MB
- Материалы: максимум 100MB
- Изображения: максимум 10MB

### Пагинация
- Максимум 100 элементов на страницу
- По умолчанию 10 элементов на страницу
