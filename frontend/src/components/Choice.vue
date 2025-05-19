<template>
  <Sidebar @filter="onFilter" />
  <div class="choice">
    <div class="choice__main">
      <ChoiceItem v-for="pet in filteredPets" :key="pet.name" :pet="pet" :openPopup="openPopup" />
    </div>

    <Popup :isOpen="isPopupOpen" :popupData="currentPet" @update:isOpen="isPopupOpen = $event" />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import ChoiceItem from './ChoiceItem.vue'
import Popup from './Popup.vue'
import Sidebar from './Sidebar.vue'

const pets = ref([
  {
    name: 'Дженнифер',
    animal: 'Собака',
    breed: 'Лабрадор',
    color: ['Золотистый', 'Белый'],
    age: '5 месяцев',
    size: 'Средний (10-20 кг)',
    sex: 'Девочка',
    compatibility: 'С детьми',
    fur: 'Средняя',
    shelter: 'Дружок',
    image: 'images/pets-jennifer-card.png',
    subtitle: 'Собака - Лабрадор',
    description: 'Описание собаки...',
    properties: {
      Возраст: '5 месяцев',
      Вакцинация: 'Пройдена',
      Болезни: 'Не указано',
      Характер: 'Дружелюбная',
    },
  },
  {
    name: 'Вуди',
    animal: 'Собака',
    breed: 'Лабрадор',
    color: 'Чёрный',
    age: '5 лет',
    size: 'Крупный (более 20 кг)',
    sex: 'Мальчик',
    compatibility: 'С собаками',
    fur: 'Длинная',
    shelter: 'Хвостатый ангел',
    image: 'images/pets-woody-card.png',
    subtitle: 'Собака - Лабрадор',
    description: 'Описание собаки...',
    properties: {
      Возраст: '5 лет',
      Вакцинация: 'Пройдена',
      Болезни: 'Не указано',
      Характер: 'Спокойная',
    },
  },
  {
    name: 'Пуговка',
    animal: 'Собака',
    breed: 'Русский той-терьер',
    color: ['Черный', 'Двухцветный', 'Смешанный'],
    age: '5 лет',
    size: 'Маленький (до 5 кг)',
    sex: 'Девочка',
    compatibility: 'Любой',
    fur: ['Гладкая', 'Короткая'],
    shelter: 'Дружок',
    image: 'images/russian-toy.jpeg',
    subtitle: 'Собака - Русский той-терьер',
    description: 'Описание собаки...',
    properties: {
      Возраст: '2 года',
      Вакцинация: 'Пройдена',
      Болезни: 'Не указано',
      Характер: 'Энергичная',
    },
  },
  {
    name: 'Автобус',
    animal: 'Собака',
    breed: 'Пекинес',
    color: ['Рыжий', 'Двухцветный'],
    age: '5 лет',
    size: 'Средний (5 - 15 кг)',
    sex: 'Мальчик',
    compatibility: 'С собаками',
    fur: 'Длинная',
    shelter: 'Хвосты и лапы',
    image: 'images/pekingese.jpg',
    subtitle: 'Собака - Пекинес',
    description: 'Описание собаки...',
    properties: {
      Возраст: '5 лет',
      Вакцинация: 'Неизвестно',
      Болезни: 'Не указано',
      Характер: 'Спокойная',
    },
  },
  {
    name: 'Джек',
    animal: 'Собака',
    breed: 'Джек-рассел-терьер',
    color: ['Белый', 'Коричневый', 'Двухцветный', 'Пятнистый'],
    age: '2 года',
    size: 'Средний (5-15 кг)',
    sex: 'Мальчик',
    compatibility: 'Любой',
    fur: 'Короткая',
    shelter: 'Хвосты и лапы',
    image: 'images/rassel.webp',
    subtitle: 'Собака - Джек-рассел-терьер',
    description: 'Описание собаки...',
    properties: {
      Возраст: '2 года',
      Вакцинация: 'Пройдена',
      Болезни: 'Не указано',
      Характер: 'Активная',
    },
  },
  {
    name: 'Хвост',
    animal: 'Собака',
    breed: 'Бигль',
    color: ['Чёрный', 'Белый', 'Коричневый', 'Трехцветный'],
    age: '10 лет',
    size: 'Средний (5-15 кг)',
    sex: 'Мальчик',
    compatibility: 'Любой',
    fur: 'Короткая',
    shelter: 'Хвостики',
    image: 'images/bigl.jpg',
    subtitle: 'Собака - Бигль',
    description: 'Описание собаки...',
    properties: {
      Возраст: '10 лет',
      Вакцинация: 'Пройдена',
      Болезни: 'Нет',
      Характер: 'Дружелюбная',
    },
  },
])

const filters = ref({
  animal: 'Любой',
  breed: 'Любая',
  color: 'Любой',
  age: 'Любой',
  size: 'Любой',
  sex: 'Любой',
  compatibility: 'Любой',
  fur: 'Любая',
  shelter: 'Любой',
})

const onFilter = (newFilters) => {
  filters.value = { ...filters.value, ...newFilters }
}

const filteredPets = computed(() => {
  const { _helpers, ...cleanFilters } = filters.value
  const isAgeInRange = filters.value._helpers?.isAgeInRange

  return pets.value.filter((pet) => {
    if (cleanFilters.animal !== 'Любой' && pet.animal !== cleanFilters.animal) {
      return false
    }

    if (cleanFilters.breed !== 'Любая' && pet.breed !== cleanFilters.breed) {
      return false
    }

    if (
      cleanFilters.color !== 'Любой' &&
      !(Array.isArray(pet.color)
        ? pet.color.includes(cleanFilters.color)
        : pet.color === cleanFilters.color)
    ) {
      return false
    }

    if (cleanFilters.age !== 'Любой' && isAgeInRange && !isAgeInRange(pet.age, cleanFilters.age)) {
      return false
    }

    if (cleanFilters.size !== 'Любой' && pet.size !== cleanFilters.size) {
      return false
    }

    if (cleanFilters.sex !== 'Любой' && pet.sex !== cleanFilters.sex) {
      return false
    }

    if (
      cleanFilters.compatibility !== 'Любой' &&
      !(pet.compatibility && pet.compatibility.includes(cleanFilters.compatibility))
    ) {
      return false
    }

    if (
      cleanFilters.fur !== 'Любая' &&
      !(Array.isArray(pet.fur) ? pet.fur.includes(cleanFilters.fur) : pet.fur === cleanFilters.fur)
    ) {
      return false
    }

    if (cleanFilters.shelter !== 'Любой' && pet.shelter !== cleanFilters.shelter) {
      return false
    }

    return true
  })
})

const isPopupOpen = ref(false)
const currentPet = ref(null)

const openPopup = (pet) => {
  currentPet.value = {
    ...pet,
    img: pet.image,
  }
  isPopupOpen.value = true
}
</script>

<style lang="scss">
.choice {
  display: flex;
  flex-direction: column;
  align-items: center;

  &__main {
    display: grid;
    grid-template-columns: repeat(3, auto);
    gap: rem(40);
    margin-bottom: rem(70);

    @include tablet {
      grid-template-columns: repeat(2, auto);
    }

    @include mobile-s {
      grid-template-columns: repeat(1, auto);
    }
  }

  &__slide {
    position: relative;
    max-width: 270px !important;
  }

  &__favorite {
    background-color: var(--color-light);
    width: rem(48);
    height: rem(48);
    position: absolute;
    top: rem(10);
    right: rem(10);
    border-radius: 50%;
    color: #f1b3b3;
    cursor: pointer;
    border: none;
  }

  &__icon {
    position: absolute;
    top: 10px;
    left: 9px;
    width: 30px;
    height: 30px;
    transition-duration: 0.5s;
  }

  &__buttons {
    display: flex;
    gap: rem(20);

    @include mobile-s {
      gap: rem(5);
    }
  }

  &__button {
    padding: rem(14) rem(13);
    background-color: transparent;
    border: rem(2) solid var(--color-accent);
    border-radius: 50%;
    width: rem(52);
    height: rem(52);
    line-height: 1;

    &:nth-child(3) {
      background-color: var(--color-accent);
    }
  }
}

.choice__main {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2.5rem;
  margin-bottom: 4.375rem;
  min-width: 873px;
}
.choice__slide {
  max-width: rem(270);
}

.slide {
  background-color: var(--color-gray-1);
  border-radius: rem(9);
  box-shadow: 0px 4px 10px -4px rgba(34, 60, 80, 0.2);
  transition-duration: 0.3s;

  &:has(.slide__button:hover) {
    transform: scale(1.1);
    transition-duration: 0.3s;
  }

  &__name {
    text-align: center;
    padding-block: rem(30);
    font-weight: 400;
  }

  &__button {
    margin: 0 auto;
    margin-bottom: rem(30);
  }
}
</style>
