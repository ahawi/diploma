<template>
  <div class="popup-overlay" v-if="isOpen" @click.self="closePopup">
    <div class="popup container">
      <button class="popup__close" @click="closePopup">×</button>
      <img class="popup__image" :src="popupData.img" :alt="popupData.name" />
      <div class="popup__text">
        <h2 class="popup__title">{{ popupData.name }}</h2>
        <p class="popup__subtitle">{{ popupData.subtitle }}</p>
        <p class="popup__description">{{ popupData.description }}</p>
        <ul class="popup__property-list">
          <li class="popup__property-item" v-for="(value, key) in popupData.properties" :key="key">
            <b>{{ key }}:</b> {{ value }}
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  isOpen: Boolean,
  popupData: Object,
})

const emit = defineEmits(['update:isOpen'])

const closePopup = () => {
  emit('update:isOpen', false)
}
</script>

<style scoped lang="scss">
.popup {
  max-width: rem(900);
  background-color: var(--color-gray-1);
  border-radius: rem(9);
  display: flex;
  box-shadow: 0px 4px 10px -4px rgba(34, 60, 80, 0.2);
  align-items: center;
  gap: rem(30);
  position: relative;
  padding: 20px;
  padding-inline: rem(50);

  &__close {
    position: absolute;
    top: -45px;
    right: -50px;
    font-size: 54px;
    line-height: 0.5;
    padding: 7px;
    border-radius: 50%;
    border: 2px solid var(--color-accent);
    color: var(--color-dark-1);

    @include mobile {
      flex-direction: column;
    }
  }

  &-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
    overflow-y: auto;
    padding: 20px;
  }

  &__image {
    width: rem(500);
    height: rem(500);

    object-fit: cover;
    object-position: center;
    display: block;

    @include tablet {
      width: rem(400);
      height: rem(400);
    }
    @include mobile {
      width: rem(270);
      height: rem(270);
    }
  }

  &__text {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: start;

    @include mobile {
      margin-block: rem(20);
    }
  }

  &__title {
    margin-bottom: rem(10);
  }

  &__subtitle {
    margin-bottom: rem(40);

    @include tablet {
      margin-bottom: rem(20);
    }
  }

  &__description {
    height: rem(100);
  }

  &__property-list {
    display: flex;
    flex-direction: column;
    gap: rem(10);
    padding-left: rem(15);
  }

  &__property {
    &-list {
      display: flex;
      flex-direction: column;
      gap: rem(10);
      padding-left: rem(15);

      @include tablet {
        gap: rem(5);
      }
    }

    &-item {
      position: relative;
      text-align: left;
      &::before {
        content: '';
        position: absolute;
        width: rem(4);
        height: rem(4);
        background-color: var(--color-accent);
        border-radius: 50%;
        top: 50%;
        left: rem(-15);
      }
    }
  }
}

.findhome__content .popup {
  height: auto;
  padding: 20px 40px;
}
</style>
