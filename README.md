# Kubernetes Monitoring with Prometheus & Grafana

Учебный проект по настройке мониторинга Kubernetes-приложения
с использованием Prometheus, Grafana и Alertmanager.

## Стек
- Kubernetes (minikube)
- Go
- Prometheus
- Grafana
- kube-prometheus-stack

## Что реализовано
- Go-приложение с `/metrics`, `/health`, `/load`
- ServiceMonitor для сбора метрик
- Grafana dashboard:
  - goroutines
  - CPU usage
  - memory usage
- Alert `DevGoAppUnavailable`:
  - срабатывает при 0 replicas
  - обрабатывает отсутствие метрик
  - auto-resolve при восстановлении

## Алерты
- DevGoAppUnavailable — **critical**
- CPU alert — запланирован на следующий этап

## Статус проекта
Проект функционально завершён.
В дальнейшем планируется расширение алертов и нагрузочное тестирование.
