FROM ubuntu:20.04
RUN useradd -m -d /home/heyue -u 2021 -U heyue
WORKDIR /home/heyue
USER heyue
ADD --chown=heyue:heyue kube /home/heyue/kube
ADD --chown=heyue:heyue template /home/heyue/template
ADD --chown=heyue:heyue autoCreateCita /home/heyue/autoCreateCita
CMD ./autoCreateCita
