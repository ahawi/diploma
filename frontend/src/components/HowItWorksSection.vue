<template>
  <section class="section">
    <div class="section__howitworks howitworks container">
      <h2 class="howitworks__title">Как это работает?</h2>
      <div class="howitworks__cards">
        <div class="howitworks__cards-item card">
          <img src="/images/pets-katrine.png" alt="" class="card__image" />
          <div class="card__description">
            <h3 class="card__title">Объединение данных из множества приютов</h3>
            <p class="card__subtitle">
              Агрегатор собирает информацию о животных из разных приютов в одном месте.
            </p>
          </div>
        </div>
        <div class="howitworks__cards-item card">
          <div class="card__description">
            <h3 class="card__title">Умный поиск и фильтрация</h3>
            <p class="card__subtitle">
              Сервис позволяет Вам настроить поиск так, чтобы найти именно того питомца.
            </p>
          </div>
          <img src="/images/pets-jennifer.png" alt="" class="card__image" />
        </div>
        <div class="howitworks__cards-item card">
          <img src="/images/pets-woody.png" alt="" class="card__image" />
          <div class="card__description">
            <h3 class="card__title">Рекомендации на основе ваших предпочтений</h3>
            <p class="card__subtitle">
              Сервис использует умные алгоритмы, чтобы предложить подходящих животных.
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup></script>

<style lang="scss" scoped>
.howitworks {
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: relative;
  z-index: 1;

  &::after {
    content: '';
    position: absolute;
    background-image: url('../../images/decor-howitworks.png');
    width: 100%;
    height: 100%;
    background-repeat: no-repeat;
    background-position: 74px 0;
    z-index: -1;
  }

  &__title {
    margin-bottom: rem(50);
    text-align: center;
  }

  &__cards {
    display: flex;
    flex-direction: column;
    gap: rem(50);

    &-item {
      display: flex;

      &:nth-child(even) {
        justify-content: end;

        @include mobile {
          justify-content: start;
        }
      }
    }
  }
}
</style>
