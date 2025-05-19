<template>
  <section class="animals-slider container">
    <h2 class="animals-slider__title">Животные, которые ищут дом</h2>

    <div class="slider-container">
      <Swiper
        :modules="[Navigation]"
        :slides-per-view="3"
        :space-between="30"
        :loop="true"
        :looped-slides="pets.length"
        :observer="true"
        :observe-parents="true"
        :navigation="{
          nextEl: '.button-next',
          prevEl: '.button-prev',
        }"
        :breakpoints="breakpoints"
      >
        <SwiperSlide v-for="pet in pets" :key="pet.name">
          <div class="slide">
            <img :src="pet.img" :alt="pet.name" class="slide__image" />
            <h4 class="slide__name">{{ pet.name }}</h4>
            <button class="slide__button button button--transparent" @click="openPopup(pet)">
              Подробнее
            </button>
          </div>
        </SwiperSlide>
      </Swiper>

      <div class="button button--arrow button-prev"></div>
      <div class="button button--arrow button-next"></div>
    </div>

    <router-link to="/animals" class="animals-slider__more button button--header">
      Подобрать друга
    </router-link>

    <!-- Попап компонент -->
    <Popup :isOpen="isPopupOpen" :popupData="currentPet" @update:isOpen="isPopupOpen = $event" />
  </section>
</template>

<script setup>
import { ref } from 'vue'
import { Swiper, SwiperSlide } from 'swiper/vue'
import { Navigation } from 'swiper/modules'
import 'swiper/css'
import 'swiper/css/navigation'
import Popup from './Popup.vue'

const breakpoints = {
  0: {
    slidesPerView: 1,
    spaceBetween: 10,
  },
  600: {
    slidesPerView: 2,
    spaceBetween: 20,
  },
  1024: {
    slidesPerView: 3,
    spaceBetween: 20,
  },
}

const pets = [
  {
    name: 'Катрин',
    img: 'images/pets-katrine-card.png',
    subtitle: 'Кошка - Русская голубая',
    description: 'Описание кошки...',
    properties: {
      Возраст: '1 год',
      Вакцинация: 'Не указано',
      Болезни: 'Не указано',
      Характер: 'Активная',
    },
  },
  {
    name: 'Дженнифер',
    img: 'images/pets-jennifer-card.png',
    subtitle: 'Собака - Лабрадор',
    description: 'Описание собаки...',
    properties: {
      Возраст: '2 года',
      Вакцинация: 'Пройдена',
      Болезни: 'Не указано',
      Характер: 'Дружелюбная',
    },
  },
  {
    name: 'Вуди',
    img: 'images/pets-woody-card.png',
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
    name: 'Катрин',
    img: 'images/pets-katrine-card.png',
    subtitle: 'Кошка - Русская голубая',
    description: 'Описание кошки...',
    properties: {
      Возраст: '1 год',
      Вакцинация: 'Не указано',
      Болезни: 'Не указано',
      Характер: 'Активная',
    },
  },
  {
    name: 'Дженнифер',
    img: 'images/pets-jennifer-card.png',
    subtitle: 'Собака - Лабрадор',
    description: 'Описание собаки...',
    properties: {
      Возраст: '2 года',
      Вакцинация: 'Пройдена',
      Болезни: 'Не указано',
      Характер: 'Дружелюбная',
    },
  },
  {
    name: 'Вуди',
    img: 'images/pets-woody-card.png',
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
    name: 'Катрин',
    img: 'images/pets-katrine-card.png',
    subtitle: 'Кошка - Русская голубая',
    description: 'Описание кошки...',
    properties: {
      Возраст: '1 год',
      Вакцинация: 'Не указано',
      Болезни: 'Не указано',
      Характер: 'Активная',
    },
  },
  {
    name: 'Дженнифер',
    img: 'images/pets-jennifer-card.png',
    subtitle: 'Собака - Лабрадор',
    description: 'Описание собаки...',
    properties: {
      Возраст: '2 года',
      Вакцинация: 'Пройдена',
      Болезни: 'Не указано',
      Характер: 'Дружелюбная',
    },
  },
  {
    name: 'Вуди',
    img: 'images/pets-woody-card.png',
    subtitle: 'Собака - Лабрадор',
    description: 'Описание собаки...',
    properties: {
      Возраст: '5 лет',
      Вакцинация: 'Пройдена',
      Болезни: 'Не указано',
      Характер: 'Спокойная',
    },
  },
]

const isPopupOpen = ref(false)
const currentPet = ref({})

const openPopup = (pet) => {
  currentPet.value = pet
  isPopupOpen.value = true
}
</script>

<style scoped lang="scss">
.animals-slider {
  padding: 2rem 0;
  text-align: center;

  &__title {
    margin-bottom: rem(40);
  }

  &__more {
    margin: rem(2) auto 0; // центрируем кнопку
    display: inline-block; // чтобы корректно работал auto

    @include hover {
      border: rem(2) solid var(--color-accent);
      color: var(--color-dark-1);
    }
  }
}

.slider-container {
  position: relative;
  max-width: 100%;
  overflow: hidden;
}

.slide {
  max-width: 270px;
  width: 100%;
  margin: 0 auto;
  background-color: var(--color-gray-1);
  border-radius: 9px;
  box-shadow: 0px 4px 10px -4px rgba(34, 60, 80, 0.2);
  padding: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: transform 0.3s;

  &:hover {
    transform: scale(1.05);
  }

  &__image {
    width: 100%;
    height: auto;
    border-radius: 6px;
  }

  &__name {
    margin: 1rem 0;
  }

  &__button {
    margin-top: auto;
  }
}

.button-prev,
.button-next {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 40px;
  height: 40px;
  background-repeat: no-repeat;
  background-position: center;
  border: 2px solid var(--color-accent);
  border-radius: 50%;
  background-color: white;
  cursor: pointer;
  z-index: 10;
}

.button-prev {
  left: 0;
  background: url('/images/arrow-left.png') no-repeat center;
}

.button-next {
  right: 0;
  background: url('/images/arrow-right.png') no-repeat center;
}

@media (max-width: 1024px) {
  .button-prev {
    left: 10px;
  }
  .button-next {
    right: 10px;
  }
}

@media (max-width: 768px) {
  .button-prev,
  .button-next {
    display: none;
  }
}

.animals-slider {
  margin-block: rem(50);
}
</style>
