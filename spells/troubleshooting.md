# troubleshooting

## definitions

- sla
- slo
- latency
- error latency
- traffic
- errors
- saturation

## golden signals

latency

cantidad de tiempo que tarda una solicitud en transferirse de un lugar a otro. Importante diferencia tiempo tanto de solicitudes exitosas como fallidas

- own 
- dependencies
- healthy 
- unhealthy

errors

tasa de solicitudes que fallan. Definir sobre cuales nos queremos alertar

- specific status codes
- under certain conditions

traffic

que tan demandado se encuentra el servicio. Diferentes formas de determinarlo.

- requests
- connections
- traffic deviations

saturation

capacidad general del servicio. CPU, threads/goroutines, hiteo de limite a otras apis, disk full, etc. no esperar al 100%

- degree of 'load' of your service 
- IO
- Disk
- RAM
- CPU
- etc

## new relic

modulo APM, application monitoring

Monitoreo del comportamiento server side. Se basa en transactions, se traducen a metricas de throughput, error rate y performance

## datadog

- informacion de la infra: CPU, Memoria
- Informacion de servicios, MySQL, KVS
- Metricas de capa de trafico: NGinx, envoys

## kibana

- permite visualizar logs aplicativos y aplicar tags

## alerts

### fixed values  

facilidad de analisis  
ejemplo: si tenemos 100 errores, alertar

### thresholds  

porcentajes de umbral preestablecidos para una alerta. No sabemos si 1000 errores es representativo asi que establecemos porcentajes de errores para el alerta
ejemplo: 10% de errores, alerta

### deviations

cuando no sabemos si 10% de error es algo malo, pero sabemos que es una desviacion de la normalidad, distinto del comportamiento comun y eso puede ser un error  
ejemplo: el comportamiento normal un lunes es 1M de solicitudes, un lunes tenemos 3M, alerta

### lack of data

falta de datos para alertar
ejemplo: varias veces el bigqueue se rompio y no alerto por falta de datos