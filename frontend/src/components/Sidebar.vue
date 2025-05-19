<template>
  <div class="sidebar">
    <form class="sidebar__filters">
      <FilterItem
        v-for="(item, index) in filters"
        :key="index"
        :label="item.label"
        :options="item.options"
        :id="item.id"
        v-model="selectedFilters[item.id]"
      />
    </form>
  </div>
</template>

<script setup>
import { ref, reactive, watch, computed } from 'vue'
import FilterItem from './FilterItem.vue'

// Исходные фильтры
const filters = ref([
  {
    label: 'Животное',
    id: 'animal',
    options: ['Любой', 'Собака'],
  },
  {
    label: 'Порода',
    id: 'breed',
    options: ['Любая', 'Бигль', 'Джек-рассел-терьер', 'Лабрадор', 'Пекинес', 'Русский той-терьер'],
  },
  {
    label: 'Цвет',
    id: 'color',
    options: [
      'Любой',
      'Черный',
      'Белый',
      'Коричневый',
      'Бежевый',
      'Серый',
      'Рыжий',
      'Золотистый',
      'Кремовый',
      'Пятнистый',
      'Двухцветный',
      'Трехцветный',
      'Тигровый',
      'Мраморный',
      'Смешанный',
    ],
  },
  {
    label: 'Возраст',
    id: 'age',
    options: ['Любой', 'До 1 года', '1–5 лет', '7–10 лет', '10 лет и старше'],
  },
  {
    label: 'Размер',
    id: 'size',
    options: ['Любой', 'Маленький (до 5 кг)', 'Средний (5 – 15 кг)', 'Крупный (более 20 кг)'],
  },
  {
    label: 'Пол',
    id: 'sex',
    options: ['Любой', 'Девочка', 'Мальчик', 'Неизвестно'],
  },
  {
    label: 'Совместим с',
    id: 'compatibility',
    options: ['Любой', 'Детьми', 'Кошками', 'Другими собаками'],
  },
  {
    label: 'Шерсть',
    id: 'fur',
    options: [
      'Любая',
      'Короткая',
      'Средняя',
      'Длинная',
      'Кудрявая',
      'Волнистая',
      'Прямая',
      'Жёсткая',
      'Гладкая',
      'Шёлковая',
      'Без шерсти',
    ],
  },
  {
    label: 'Приют',
    id: 'shelter',
    options: ['Любой', 'Дружок', 'Хвостатый ангел', 'Хвосты и лапы', 'Хвостики'],
  },
])

const selectedFilters = reactive({})

// Инициализация выбранных значений
filters.value.forEach((filter) => {
  selectedFilters[filter.id] = filter.options[0]
})

const parseAge = (ageStr) => {
  const lower = ageStr.toLowerCase().trim()
  if (lower.includes('месяц') || lower.includes('до 1 года')) return 0.5
  const match = lower.match(/(\d+)\s*(год|лет)/)
  return match ? parseInt(match[1]) : null
}

// Проверка попадания в диапазон возраста
const isAgeInRange = (petAge, filterRange) => {
  const ageInYears = parseAge(petAge)
  if (ageInYears === null) return false

  switch (filterRange) {
    case 'До 1 года':
      return ageInYears < 1
    case '1–5 лет':
      return ageInYears >= 1 && ageInYears <= 5
    case '7–10 лет':
      return ageInYears >= 7 && ageInYears <= 10
    case '10 лет и старше':
      return ageInYears >= 10
    default:
      return true
  }
}

// Эмит обновлённых значений
const emit = defineEmits(['filter'])

watch(
  selectedFilters,
  () => {
    emit('filter', {
      ...selectedFilters,
      _helpers: {
        isAgeInRange, // передаем помощник для фильтрации по возрасту
      },
    })
  },
  { deep: true },
)
</script>

<style lang="scss">
.sidebar {
  &__filters {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: rem(20);
    background-color: var(--color-gray-1);
    padding: rem(34) rem(20);
    border-radius: rem(9);
    width: rem(272);
    box-shadow: 0px 4px 10px -4px rgba(34, 60, 80, 0.2);

    @include laptop {
      max-width: rem(250);
    }

    @include tablet {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      width: 100%;
      max-width: 100%;
    }

    @include mobile-s {
      grid-template-columns: repeat(1, 1fr);
    }
  }

  &__item {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: rem(10);
  }

  &__title {
    text-align: center;
    font-style: italic;
  }

  &__select {
    appearance: none;
    border: rem(2) solid var(--color-accent);
    border-radius: rem(100);
    padding: rem(10) rem(18);
    outline: none;
    color: #767676;
    background-color: transparent;
    background-image: url('../../images/select-arrow.svg');
    background-repeat: no-repeat;
    background-position: right rem(18) center;
    padding-right: rem(40);
  }

  .slide__name {
    @include laptop {
      font-size: rem(26);
    }
  }
}
</style>
