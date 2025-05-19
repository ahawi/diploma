<template>
  <section class="section">
    <div class="section__about about container" id="about">
      <img src="/images/about-pets.png" alt="Животные" class="about__image" />
      <div class="about__description">
        <h2 class="about__title">О сервисе PawfectMatch</h2>
        <p class="about__subtitle">
          Этот сервис создан для того, чтобы помочь животным из приютов найти любящих хозяев, а
          людям — обрести верного друга. Он работает как умный помощник, который объединяет
          информацию о животных из разных приютов и использует современные технологии, чтобы сделать
          процесс поиска и подбора максимально удобным и эффективным.
        </p>
      </div>
    </div>
  </section>
</template>

<script setup></script>

<style lang="scss" scoped>
.about {
  display: flex;
  align-items: center;
  justify-content: center;
  column-gap: rem(120);
  position: relative;

  @include tablet {
    flex-direction: column-reverse;
    row-gap: rem(30);
  }

  &::after {
    content: '';
    position: absolute;
    background-image: url('../../images/stars-about.png');
    width: 100%;
    height: 100%;
    background-repeat: no-repeat;
    background-position: 60px 75px;
    z-index: -1;
  }

  &__description {
    display: flex;
    flex-direction: column;
    row-gap: rem(70);
    max-width: rem(430);

    @include tablet {
      row-gap: rem(30);
      max-width: rem(600);
    }
  }

  &__image {
    max-width: 80%;
  }
}
</style>
